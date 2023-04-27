package logs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

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

func (s *service) ComponentLog(appName, envName, componentName string, options *ComponentLogOptions) (io.Reader, error) {
	if options == nil {
		options = &ComponentLogOptions{}
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

func (s *service) ComponentPodInventory(appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error) {
	if options == nil {
		options = &ComponentPodInventoryOptions{}
	}

	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{SuppressWarning: true})).
		UnsafeAdd(componentPodInventory).
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

type logReader struct {
	source *azquery.Table
	row    int
	offset int
	logCol int
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

func (r *logReader) Read(p []byte) (n int, err error) {
	if r.source == nil {
		return 0, io.EOF
	}

	bufLen := len(p)
	var bpos int

	for bpos < bufLen {
		cp, err := r.copyRow(p[bpos:])
		if err != nil {
			return bpos, err
		}
		bpos += cp

	}

	return bpos, nil
}

func (r *logReader) copyRow(p []byte) (int, error) {
	rowCount := len(r.source.Rows)
	if r.row >= rowCount {
		return 0, io.EOF
	}
	currRow, ok := r.source.Rows[r.row][r.logCol].(string)
	if !ok {
		return 0, errors.New("unexpected data in log")
	}
	currRow += "\n"
	cp := copy(p, currRow[r.offset:])
	r.offset += cp

	if r.offset >= len(currRow) {
		r.row++
		r.offset = 0
	}
	return cp, nil
}
