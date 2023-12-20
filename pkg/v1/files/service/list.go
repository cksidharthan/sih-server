package service

import (
	"context"
	domain "github.com/cksidharthan/sih-server/pkg/v1/files"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// List lists the files in the server
func (s *FilesService) List(ctx context.Context) ([]domain.SihFile, error) {
	var files []domain.SihFile

	err := filepath.Walk(s.envCfg.UploadFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, s.ConvertFileNameToSihFile(info.Name()))
		}
		return nil
	})

	return files, err
}

func (s *FilesService) ConvertFileNameToSihFile(fileName string) domain.SihFile {
	fileNameSplit := strings.Split(fileName, "__++")
	// get the file size
	fileInfo, err := os.Stat(s.envCfg.UploadFolder + "/" + fileName)
	if err != nil {
		logrus.Error(err)
	}
	fileSize := fileInfo.Size()

	return domain.SihFile{
		UUID:      fileNameSplit[0],
		Timestamp: fileNameSplit[1],
		FileName:  fileNameSplit[2],
		Size:      fileSize,
	}
}
