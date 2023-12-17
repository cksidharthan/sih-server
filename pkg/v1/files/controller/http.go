package controller

import (
	"github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type FileHandler struct {
	fx.In

	RoutesV1 *gin.RouterGroup `name:"v1"`
	Service  *service.FilesService
}

func New(u FileHandler) {
	router := u.RoutesV1.Group("/files")
	router.POST("/upload", UploadFile(u.Service))
	router.GET("/list", ListFiles(u.Service))
	router.GET("/download", Download(u.Service))
}
