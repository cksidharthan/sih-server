package controller

import (
	"github.com/cksidharthan/sih-server/pkg/v1/upload/service"
	"github.com/gin-gonic/gin"
)

func UploadFile(service service.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
