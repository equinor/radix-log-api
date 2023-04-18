package uriparams

type Component struct {
	ComponentName string `uri:"componentName" binding:"required"`
}
