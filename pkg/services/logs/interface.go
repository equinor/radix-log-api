package logs

import (
	"context"
	"io"
)

type Service interface {
	ComponentLog(ctx context.Context, appName, envName, componentName string, options *LogOptions) (io.Reader, error)
	ComponentPodLog(ctx context.Context, appName, envName, componentName, replicaName string, options *LogOptions) (io.Reader, error)
	ComponentContainerLog(ctx context.Context, appName, envName, componentName, replicaName, containerId string, options *LogOptions) (io.Reader, error)
	ComponentInventory(ctx context.Context, appName, envName, componentName string, options *ComponentPodInventoryOptions) ([]Pod, error)
}
