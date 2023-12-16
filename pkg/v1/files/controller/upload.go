package controller

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"github.com/gin-gonic/gin"
)

// UploadFile is a gin handler function that uploads a file to the server
func UploadFile(service *service.FilesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "file not found",
			})
			return
		}

		// save the file to disk
		err = service.Upload(c, file)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error saving file",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	}
}
