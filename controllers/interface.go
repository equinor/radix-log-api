package controllers

import "github.com/gin-gonic/gin"

type Endpoint struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type Controller interface {
	Endpoints() []Endpoint
}
