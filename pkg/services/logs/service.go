package logs

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/data/types"
	"github.com/Azure/azure-kusto-go/kusto/kql"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/equinor/radix-log-api/pkg/aztable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ContainerLogType string

const (
	ContainerLogTypeV1   ContainerLogType = "ContainerLog"
	ContainerLogTypeV2   ContainerLogType = "ContainerLogV2"
	ContainerLogTypeBoth ContainerLogType = "Both"
)

type service struct {
	logsClient        *azquery.LogsClient
	workspaceId       string
	containerLogQuery string
}

func New(logsClient *azquery.LogsClient, workspaceId string, logType ContainerLogType) Service {
	return &service{
		logsClient:        logsClient,
		workspaceId:       workspaceId,
		containerLogQuery: getContainerJoinQuery(logType),
	}
}

// TODO: Remove this when legacy ContainerLog is no longer in use
func getContainerJoinQuery(logType ContainerLogType) string {
	var containerLogQuery string
	switch logType {
	case ContainerLogTypeV1:
		log.Info().Str("table", "ContainerLog").Msg("Configuring Log Analytics")
		containerLogQuery = joinContainerLogV1
	case ContainerLogTypeV2:
		log.Info().Str("table", "ContainerLogV2").Msg("Configuring Log Analytics")
		containerLogQuery = joinContainerLogV2
	case ContainerLogTypeBoth:
		log.Info().Str("table", "ContainerLog,ContainerLogV2").Msgf("Configuring Log Analytics")
		containerLogQuery = joinContainerBoth
	default:
		log.Warn().Str("table", "ContainerLog").Msg("Configuring Log Analytics, unknown Log Type, fallback to default")
		containerLogQuery = joinContainerLogV1
	}
	return containerLogQuery
}

func (s *service) ComponentLog(ctx context.Context, appName, appId, envName, componentName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:         kusto.ParamType{Type: types.String, Default: appId},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentPodLog(ctx context.Context, appName, appId, envName, componentName, podName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:         kusto.ParamType{Type: types.String, Default: appId},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
			paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentPodLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentContainerLog(ctx context.Context, appName, appId, envName, componentName, podName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:         kusto.ParamType{Type: types.String, Default: appId},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
			paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
			paramContainerId:   kusto.ParamType{Type: types.String, Default: containerId},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentContainerLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentInventory(ctx context.Context, appName, appId, envName, componentName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:         kusto.ParamType{Type: types.String, Default: appId},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) JobInventory(ctx context.Context, appName, appId, envName, jobComponentName, jobName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:            kusto.ParamType{Type: types.String, Default: appId},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) JobLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:            kusto.ParamType{Type: types.String, Default: appId},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) JobPodLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName, replicaName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:            kusto.ParamType{Type: types.String, Default: appId},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
			paramPodName:          kusto.ParamType{Type: types.String, Default: replicaName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobPodLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) JobContainerLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramAppId:            kusto.ParamType{Type: types.String, Default: appId},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
			paramPodName:          kusto.ParamType{Type: types.String, Default: replicaName},
			paramContainerId:      kusto.ParamType{Type: types.String, Default: containerId},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobContainerLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) PipelineJobInventory(ctx context.Context, appName, appId, pipelineJobName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:       kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-app", appName)},
			paramPipelineJobName: kusto.ParamType{Type: types.String, Default: pipelineJobName},
			paramAppId:           kusto.ParamType{Type: types.String, Default: appId},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getPipelineJobInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) PipelineJobContainerLog(ctx context.Context, appName, appId, pipelineJobName string, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:       kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-app", appName)},
			paramPodName:         kusto.ParamType{Type: types.String, Default: replicaName},
			paramContainerId:     kusto.ParamType{Type: types.String, Default: containerId},
			paramAppId:           kusto.ParamType{Type: types.String, Default: appId},
			paramPipelineJobName: kusto.ParamType{Type: types.String, Default: pipelineJobName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getPipelineJobContainerLogQuery(s.containerLogQuery))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) executeInventoryQuery(ctx context.Context, builder *kql.Builder, options *InventoryOptions) ([]Pod, error) {
	logger := zerolog.Ctx(ctx)
	debugEvent := logger.Debug()

	if options == nil {
		options = &InventoryOptions{}
	}

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
		debugEvent.Time("start", options.Timeinterval.Start).Time("end", options.Timeinterval.End)
	}

	query := builder.String()
	debugEvent.Str("query", query).Msg("Execute inventory query")
	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		logger.Warn().Err(resp.Error).Msg("Inventory query returned a warning")
	}

	reader := newTableReader(resp.Tables[0])
	podmap := slice.Reduce(resp.Tables[0].Rows, map[string]*Pod{}, func(acc map[string]*Pod, row azquery.Row) map[string]*Pod {
		podName := reader.Value(row, "Name", "").(string)
		lastTimeGenerated := mustParseTime(reader.Value(row, "LastTimeGenerated", "").(string))
		pod, ok := acc[podName]
		if !ok {
			pod = &Pod{
				Name:              podName,
				CreationTimestamp: mustParseTime(reader.Value(row, "PodCreationTimeStamp", "").(string)),
				LastKnown:         lastTimeGenerated,
				Containers:        []Container{},
			}
		}
		if lastTimeGenerated.After(pod.LastKnown) {
			pod.LastKnown = lastTimeGenerated
		}
		pod.Containers = append(pod.Containers,
			Container{
				Id:                reader.Value(row, "ContainerID", "").(string),
				Name:              reader.Value(row, "ContainerNameShort", "").(string),
				LastKnown:         lastTimeGenerated,
				CreationTimestamp: mustParseTime(reader.Value(row, "ContainerCreationTimeStamp", "").(string)),
			})
		acc[podName] = pod
		return acc
	})

	pods := make([]Pod, 0, len(podmap))
	for _, pod := range podmap {
		pods = append(pods, *pod)
	}
	return pods, nil
}

func (s *service) executeLogQuery(ctx context.Context, builder *kql.Builder, options *LogOptions) (io.Reader, error) {
	logger := zerolog.Ctx(ctx)
	debugEvent := logger.Debug()

	if options == nil {
		options = &LogOptions{}
	}

	if options.LimitRows != nil {
		builder = builder.AddLiteral("| take ").AddInt(int32(*options.LimitRows))
		debugEvent.Int("limit_rows", *options.LimitRows)
	}

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
		debugEvent.Time("start", options.Timeinterval.Start).Time("end", options.Timeinterval.End)
	}

	query := builder.String()
	debugEvent.Str("query", query).Msg("Execute log query")

	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}
	// TODO: Max response from log analytics is 64MB. Result is truncated if limit is exceeded.
	// When result is trucated, the reason is described in resp.Error.
	// Tried `set notruncation` (https://aka.ms/kustoquerylimits) but it had no effect.
	// For now we just log a warning if response contain an error, but perhaps we must forward the
	// error to the user.
	if resp.Error != nil {
		logger.Warn().Err(resp.Error).Msg("Log query returned a warning")
	}

	return aztable.NewReader(resp.Tables[0], 1), nil
}

func mustParseTime(t string) time.Time {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}
