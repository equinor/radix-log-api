package logs

import "io"

type Interface interface {
	GetLogs(appName, envName, componentName string, options *GetLogsQueryOptions) (io.Reader, error)
}
