package controller

import (
	"github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"github.com/gin-gonic/gin"
)

func ListFiles(service *service.FilesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// save the file to disk
		files, err := service.List(c)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error listing files",
			})
			return
		}

		c.JSON(200, gin.H{
			"data": files,
		})
	}
}
