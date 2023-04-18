package uriparams

type App struct {
	AppName string `uri:"appName" binding:"required"`
}
