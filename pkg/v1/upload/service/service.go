package service

import (
	"github.com/cksidharthan/sih-server/pkg/config"
	"github.com/sirupsen/logrus"
)

type UploadService struct {
	log    *logrus.Entry
	envCfg *config.Config
}

func New(log *logrus.Logger, envCfg *config.Config) *UploadService {
	return &UploadService{
		log:    log.WithField("pkg", "upload"),
		envCfg: envCfg,
	}
}
