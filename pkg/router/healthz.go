package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultEndpoints - creates the health endpoint that will be exposed by the application when it is started.
func DefaultEndpoints(engine *gin.Engine) {
	// Healthz endpoint
	engine.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})
}
