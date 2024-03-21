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
)

type ContainerLogType int

const (
	ContainerLogTypeV1 ContainerLogType = iota
	ContainerLogTypeV2
	ContainerLogTypeBoth
)

type service struct {
	logsClient           *azquery.LogsClient
	workspaceId          string
	containerLogJoinName string
}

func New(logsClient *azquery.LogsClient, workspaceId string, logType ContainerLogType) Service {
	var containerLogJoinName string
	switch logType {
	case ContainerLogTypeV1:
		containerLogJoinName = joinContainerLogV1
	case ContainerLogTypeV2:
		containerLogJoinName = joinContainerLogV2
	case ContainerLogTypeBoth:
		containerLogJoinName = joinContainerBoth
	}

	return &service{
		logsClient:           logsClient,
		workspaceId:          workspaceId,
		containerLogJoinName: containerLogJoinName,
	}
}

func (s *service) ComponentLog(ctx context.Context, appName, envName, componentName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentPodLog(ctx context.Context, appName, envName, componentName, podName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
			paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentPodLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentContainerLog(ctx context.Context, appName, envName, componentName, podName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
			paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
			paramContainerId:   kusto.ParamType{Type: types.String, Default: containerId},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentContainerLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) ComponentInventory(ctx context.Context, appName, envName, componentName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
			paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getComponentInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) JobInventory(ctx context.Context, appName, envName, jobComponentName, jobName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) JobLog(ctx context.Context, appName, envName, jobComponentName, jobName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) JobPodLog(ctx context.Context, appName, envName, jobComponentName, jobName, replicaName string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
			paramPodName:          kusto.ParamType{Type: types.String, Default: replicaName},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobPodLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) JobContainerLog(ctx context.Context, appName, envName, jobComponentName, jobName, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:        kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
			paramAppName:          kusto.ParamType{Type: types.String, Default: appName},
			paramJobComponentName: kusto.ParamType{Type: types.String, Default: jobComponentName},
			paramJobName:          kusto.ParamType{Type: types.String, Default: jobName},
			paramPodName:          kusto.ParamType{Type: types.String, Default: replicaName},
			paramContainerId:      kusto.ParamType{Type: types.String, Default: containerId},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getJobContainerLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) PipelineJobInventory(ctx context.Context, appName, pipelineJobName string, options *InventoryOptions) ([]Pod, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:       kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-app", appName)},
			paramPipelineJobName: kusto.ParamType{Type: types.String, Default: pipelineJobName},
		},
	)
	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getPipelineJobInventoryQuery())

	return s.executeInventoryQuery(ctx, builder, options)
}

func (s *service) PipelineJobContainerLog(ctx context.Context, appName, pipelineJobName string, replicaName, containerId string, options *LogOptions) (io.Reader, error) {
	params := kusto.NewDefinitions().Must(
		kusto.ParamTypes{
			paramNamespace:       kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-app", appName)},
			paramAppName:         kusto.ParamType{Type: types.String, Default: appName},
			paramPipelineJobName: kusto.ParamType{Type: types.String, Default: pipelineJobName},
			paramPodName:         kusto.ParamType{Type: types.String, Default: replicaName},
			paramContainerId:     kusto.ParamType{Type: types.String, Default: containerId},
		},
	)

	builder := kql.New("").
		AddUnsafe(params.String()).
		AddUnsafe(getPipelineJobContainerLogQuery(joinContainerLogV1))

	return s.executeLogQuery(ctx, builder, options)
}

func (s *service) executeInventoryQuery(ctx context.Context, builder *kql.Builder, options *InventoryOptions) ([]Pod, error) {
	if options == nil {
		options = &InventoryOptions{}
	}

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
	}

	query := builder.String()
	zerolog.Ctx(ctx).Debug().Str("query", query).Msg("Execute inventory query")
	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
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
	if options == nil {
		options = &LogOptions{}
	}

	if options.LimitRows != nil {
		builder = builder.AddLiteral("| take ").AddInt(int32(*options.LimitRows))
	}

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
	}

	query := builder.String()
	zerolog.Ctx(ctx).Debug().Str("query", query).Msg("Execute log query")
	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}

	return aztable.NewReader(resp.Results.Tables[0], 3), nil
}

func mustParseTime(t string) time.Time {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}
