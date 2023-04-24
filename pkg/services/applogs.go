package services

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
	"github.com/equinor/radix-common/utils/pointers"
)

const (
	paramNamespace     = "ParamNamespace"
	paramAppName       = "ParamAppName"
	paramComponentName = "ParamComponentName"
)

var (
	logQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s
	| where ContainerID != ""
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s
	| summarize ContainerCreationTimeStamp=min(ContainerCreationTimeStamp) by ContainerID, Name
	| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated asc`, paramNamespace, paramAppName, paramComponentName)
)

var (
	defaultGetLogsQueryOptions *GetLogsQueryOptions = &GetLogsQueryOptions{
		LimitRows: pointers.Ptr(1000),
	}
)

type GetLogsQueryOptions struct {
	LimitRows *int
}

type AppLogs interface {
	GetLogs(appName, envName, componentName string, options *GetLogsQueryOptions) (io.Reader, error)
}

type appLogs struct {
	logsClient  *azquery.LogsClient
	workspaceId string
}

func (s *appLogs) GetLogs(appName, envName, componentName string, options *GetLogsQueryOptions) (io.Reader, error) {
	if options == nil {
		options = defaultGetLogsQueryOptions
	}

	kql := kusto.NewStmt("", kusto.UnsafeStmt(unsafe.Stmt{})).
		UnsafeAdd(logQuery).
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

	query := kql.String()
	timspan := azquery.NewTimeInterval(time.Now().Add(-48*time.Hour), time.Now())
	resp, err := s.logsClient.QueryWorkspace(context.TODO(), s.workspaceId, azquery.Body{Query: &query, Timespan: &timspan}, nil)
	if err != nil {
		return nil, err
	}

	return &logReader{source: resp.Results.Tables[0], logCol: 3}, nil
}

func NewAppLogs(logsClient *azquery.LogsClient, workspaceId string) AppLogs {
	return &appLogs{
		logsClient:  logsClient,
		workspaceId: workspaceId,
	}
}

type logReader struct {
	source *azquery.Table
	row    int
	offset int
	logCol int
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
