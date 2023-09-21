package params

type Job struct {
	JobName string `uri:"jobName" binding:"required"`
}
