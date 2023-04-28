package logs

import (
	"context"
	"fmt"
	"io"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/data/types"
	"github.com/Azure/azure-kusto-go/kusto/unsafe"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/equinor/radix-common/utils/slice"
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
	if options == nil {
		options = &LogOptions{}
	}

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

	if options.LimitRows != nil {
		kql.UnsafeAdd(fmt.Sprintf("| take %d", *options.LimitRows))
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

	return &logReader{source: resp.Results.Tables[0], logCol: 3}, nil
}

func (s *service) ComponentPodLog(appName, envName, componentName, podName string, options *LogOptions) (io.Reader, error) {
	return nil, nil
}
func (s *service) ComponentContainerLog(appName, envName, componentName, podName, containerId string, options *LogOptions) (io.Reader, error) {
	return nil, nil
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

	podmap := slice.Reduce(resp.Tables[0].Rows, map[string]*Pod{}, func(acc map[string]*Pod, row azquery.Row) map[string]*Pod {
		podName := row[0].(string)
		pod, ok := acc[podName]
		if !ok {
			pod = &Pod{Name: podName, CreationTimestamp: mustParseTime(row[2].(string)), Containers: []Container{}}
		}
		pod.Containers = append(pod.Containers, Container{Id: row[1].(string), CreationTimestamp: mustParseTime(row[3].(string))})
		acc[podName] = pod
		return acc
	})

	pods := make([]Pod, 0, len(podmap))
	for _, pod := range podmap {
		pods = append(pods, *pod)
	}
	return pods, nil
}
