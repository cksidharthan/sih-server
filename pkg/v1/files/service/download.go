package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (s *FilesService) Download(ctx context.Context, fileUUID string) (string, error) {
	var filename string

	// Open the directory
	err := filepath.Walk(s.envCfg.UploadFolder, func(path string, info os.FileInfo, err error) error {
		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if the file name has the specified prefix
		if strings.HasPrefix(info.Name(), fileUUID) {
			fmt.Printf("File with prefix %s found: %s\n", fileUUID, info.Name())
			filename = s.envCfg.UploadFolder + "/" + info.Name()
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return filename, nil
}
