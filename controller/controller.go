package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	// Register method can received Gin Engine. Gin-Gonic is micro-framework for web applications.
	// We can manage web requests with Gin-Gonic.
	Register(engine *gin.Engine)
}
