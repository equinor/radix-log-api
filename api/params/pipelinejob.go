package params

type PipelineJob struct {
	PipelineJobName string `uri:"pipelineJobName" binding:"required"`
}
