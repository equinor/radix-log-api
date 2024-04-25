package main

import (
	"github.com/equinor/radix-log-api/cmd"
	_ "github.com/equinor/radix-log-api/docs"
)

// @BasePath /api/v1
// @Schemes http https
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					Bearer token is currently not supported by go-swag. Use "Bearer <JWT>" in value.
func main() {
	cmd.Run()
}
