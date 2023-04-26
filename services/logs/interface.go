package logs

import "io"

type Service interface {
	Component(appName, envName, componentName string, options *GetLogsQueryOptions) (io.Reader, error)
}
