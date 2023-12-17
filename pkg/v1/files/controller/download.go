package controller

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"github.com/gin-gonic/gin"
)

func Download(service *service.FilesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the filename from the parameter
		filename := c.Query("fileName")
		if filename == "" {
			c.JSON(400, gin.H{
				"message": "filename not found",
			})
			return
		}

		fmt.Println(filename)

		// save the file to disk
		file, err := service.Download(c, filename)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error listing files",
			})
			return
		}

		c.File(file)
	}
}
