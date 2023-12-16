package controller

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/v1/upload/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func UploadFile(service *service.UploadService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"message": "file not found",
			})
			return
		}

		// Save the uploaded file to a temporary location
		uploadedFilePath := "temp/" + file.Filename
		if err := c.SaveUploadedFile(file, uploadedFilePath); err != nil {
			c.JSON(500, gin.H{
				"message": "failed to save file",
			})
			return
		}
		defer func() {
			// Remove the temporary file after processing
			if err := os.Remove(uploadedFilePath); err != nil {
				fmt.Println("Error removing temporary file:", err)
			}
		}()

		// Read the content of the file
		content, err := ioutil.ReadFile(uploadedFilePath)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to read file",
			})
			return
		}

		// Print the content of the file
		fmt.Println("File Content:")
		fmt.Println(string(content))

		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
