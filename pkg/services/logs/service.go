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
)

type service struct {
	logsClient  *azquery.LogsClient
	workspaceId string
}

func New(logsClient *azquery.LogsClient, workspaceId string) Service {
	return &service{
		logsClient:  logsClient,
		workspaceId: workspaceId,
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
		AddUnsafe(componentLogQuery)

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
		AddUnsafe(componentPodLogQuery)

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
		AddUnsafe(componentContainerLogQuery)

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
		AddUnsafe(componentInventoryQuery)

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
		AddUnsafe(jobInventoryQuery)

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
		AddUnsafe(jobLogQuery)

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
		AddUnsafe(jobPodLogQuery)

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
		AddUnsafe(jobContainerLogQuery)

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
	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}

	colMap := map[string]int{}
	for i, col := range resp.Tables[0].Columns {
		colMap[*col.Name] = i
	}

	podmap := slice.Reduce(resp.Tables[0].Rows, map[string]*Pod{}, func(acc map[string]*Pod, row azquery.Row) map[string]*Pod {
		podName := row[colMap["Name"]].(string)
		lastTimeGenerated := mustParseTime(row[colMap["LastTimeGenerated"]].(string))
		pod, ok := acc[podName]
		if !ok {
			pod = &Pod{
				Name:              podName,
				CreationTimestamp: mustParseTime(row[colMap["PodCreationTimeStamp"]].(string)),
				LastKnown:         lastTimeGenerated,
				Containers:        []Container{},
			}
		}
		if lastTimeGenerated.After(pod.LastKnown) {
			pod.LastKnown = lastTimeGenerated
		}
		pod.Containers = append(pod.Containers,
			Container{
				Id:                row[colMap["ContainerID"]].(string),
				LastKnown:         lastTimeGenerated,
				CreationTimestamp: mustParseTime(row[colMap["ContainerCreationTimeStamp"]].(string)),
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
	resp, err := s.logsClient.QueryWorkspace(ctx, s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}

	return aztable.NewReader(resp.Results.Tables[0], 3), nil
}

func mustParseTime(t string) time.Time {
	if t == "" {
		fmt.Println("")
	}
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}
