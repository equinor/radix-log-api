package controllers

import "github.com/gin-gonic/gin"

type Endpoint struct {
	Path                  string
	Method                string
	Handler               gin.HandlerFunc
	AuthorizationPolicies []string
}

// Controller defines the endpoints for a HTTP API controller.
type Controller interface {
	// Endpoints returns the endpoints defined for the controller.
	Endpoints() []Endpoint
}
