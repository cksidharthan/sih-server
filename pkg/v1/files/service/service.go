package service

import (
	"github.com/cksidharthan/sih-server/pkg/config"
	"github.com/sirupsen/logrus"
)

type FilesService struct {
	log    *logrus.Entry
	envCfg *config.Config
}

func New(log *logrus.Logger, envCfg *config.Config) *FilesService {
	return &FilesService{
		log:    log.WithField("pkg", "upload"),
		envCfg: envCfg,
	}
}
