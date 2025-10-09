package logs

import (
	"context"
	"io"
)

type Service interface {
	ComponentLog(ctx context.Context, appName, appId, envName, componentName string, options *LogOptions) (io.Reader, error)
	ComponentPodLog(ctx context.Context, appName, appId, envName, componentName, replicaName string, options *LogOptions) (io.Reader, error)
	ComponentContainerLog(ctx context.Context, appName, appId, envName, componentName, replicaName, containerId string, options *LogOptions) (io.Reader, error)
	ComponentInventory(ctx context.Context, appName, appId, envName, componentName string, options *InventoryOptions) ([]Pod, error)
	JobInventory(ctx context.Context, appName, appId, envName, jobComponentName, jobName string, options *InventoryOptions) ([]Pod, error)
	JobLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName string, options *LogOptions) (io.Reader, error)
	JobPodLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName, replicaName string, options *LogOptions) (io.Reader, error)
	JobContainerLog(ctx context.Context, appName, appId, envName, jobComponentName, jobName, replicaName, containerId string, options *LogOptions) (io.Reader, error)
	PipelineJobInventory(ctx context.Context, appName, appId, pipelineJobName string, options *InventoryOptions) ([]Pod, error)
	PipelineJobContainerLog(ctx context.Context, appName, appId, pipelineJobName string, replicaName, containerId string, options *LogOptions) (io.Reader, error)
}
