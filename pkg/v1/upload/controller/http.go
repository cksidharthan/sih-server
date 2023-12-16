package controller

import (
	"github.com/cksidharthan/sih-server/pkg/v1/upload/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UploadHandler struct {
	fx.In

	RoutesV1 *gin.RouterGroup `name:"v1"`
	Service  service.UploadService
}

func New(u UploadHandler) {
	router := u.RoutesV1.Group("/download")
	router.POST("", UploadFile(u.Service))
}
