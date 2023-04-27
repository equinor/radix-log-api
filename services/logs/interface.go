package logs

import "io"

type Service interface {
	ComponentLog(appName, envName, componentName string, options *ComponentLogOptions) (io.Reader, error)
	ComponentPodInventory(appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error)
}
