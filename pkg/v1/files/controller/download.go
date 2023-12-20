package controller

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Download(service *service.FilesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the filename from the parameter
		filename := c.Query("uuid")
		fmt.Println(filename)
		if filename == "" {
			c.JSON(400, gin.H{
				"message": "filename not found",
			})
			return
		}

		// save the file to disk
		file, err := service.Download(c, filename)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error listing files",
			})
			return
		}

		// Set the "Content-Disposition" header for the browser to prompt download
		c.Header("Content-Disposition", "attachment; filename="+file)
		c.Header("Content-Type", "application/text/plain")

		// read the file and convert it to bytes
		// Read the entire file into a byte slice
		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		c.Writer.Write(fileContent)
		c.JSON(http.StatusOK, gin.H{
			"msg": "Download file successfully",
		})
	}
}
