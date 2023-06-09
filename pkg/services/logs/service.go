package logs

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/data/types"
	"github.com/Azure/azure-kusto-go/kusto/unsafe"
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

func (s *service) ComponentLog(appName, envName, componentName string, options *LogOptions) (io.Reader, error) {
	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{SuppressWarning: true})).
		UnsafeAdd(componentLogQuery).
		MustDefinitions(
			kusto.NewDefinitions().Must(
				kusto.ParamTypes{
					paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
					paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
					paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
				},
			),
		)

	return s.executeLogQuery(kql, options)
}

func (s *service) ComponentPodLog(appName, envName, componentName, podName string, options *LogOptions) (io.Reader, error) {
	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{SuppressWarning: true})).
		UnsafeAdd(componentPodLogQuery).
		MustDefinitions(
			kusto.NewDefinitions().Must(
				kusto.ParamTypes{
					paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
					paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
					paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
					paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
				},
			),
		)

	return s.executeLogQuery(kql, options)
}
func (s *service) ComponentContainerLog(appName, envName, componentName, podName, containerId string, options *LogOptions) (io.Reader, error) {
	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{SuppressWarning: true})).
		UnsafeAdd(componentContainerLogQuery).
		MustDefinitions(
			kusto.NewDefinitions().Must(
				kusto.ParamTypes{
					paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
					paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
					paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
					paramPodName:       kusto.ParamType{Type: types.String, Default: podName},
					paramContainerId:   kusto.ParamType{Type: types.String, Default: containerId},
				},
			),
		)

	return s.executeLogQuery(kql, options)
}

func (s *service) executeLogQuery(kql kusto.Stmt, options *LogOptions) (io.Reader, error) {
	if options == nil {
		options = &LogOptions{}
	}

	if options.LimitRows != nil {
		kql = kql.UnsafeAdd(fmt.Sprintf("| take %d", *options.LimitRows))
	}

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
	}

	query := kql.String()
	resp, err := s.logsClient.QueryWorkspace(context.TODO(), s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}

	return aztable.NewReader(resp.Results.Tables[0], 3), nil
}

func (s *service) ComponentInventory(appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error) {
	if options == nil {
		options = &ComponentPodInventoryOptions{}
	}

	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{SuppressWarning: true})).
		UnsafeAdd(componentInventory).
		MustDefinitions(
			kusto.NewDefinitions().Must(
				kusto.ParamTypes{
					paramNamespace:     kusto.ParamType{Type: types.String, Default: fmt.Sprintf("%s-%s", appName, envName)},
					paramAppName:       kusto.ParamType{Type: types.String, Default: appName},
					paramComponentName: kusto.ParamType{Type: types.String, Default: componentName},
				},
			),
		)

	timspan := azquery.TimeInterval("")
	if options.Timeinterval != nil {
		timspan = options.Timeinterval.AzQueryTimeinterval()
	}

	query := kql.String()
	resp, err := s.logsClient.QueryWorkspace(context.TODO(), s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
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
