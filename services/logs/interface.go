package logs

import "io"

type Service interface {
	ComponentLog(appName, envName, componentName string, options *LogOptions) (io.Reader, error)
	ComponentPodLog(appName, envName, componentName, replicaName string, options *LogOptions) (io.Reader, error)
	ComponentContainerLog(appName, envName, componentName, replicaName, containerId string, options *LogOptions) (io.Reader, error)
	ComponentInventory(appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error)
}
