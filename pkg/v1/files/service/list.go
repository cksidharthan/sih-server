package service

import (
	"context"
	"os"
	"path/filepath"
)

// List lists the files in the server
func (s *FilesService) List(ctx context.Context) ([]string, error) {
	var files []string

	err := filepath.Walk(s.envCfg.UploadFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
