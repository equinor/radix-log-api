package params

type Env struct {
	EnvName string `uri:"envName" binding:"required"`
}
