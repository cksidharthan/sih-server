package service

import (
	"context"
)

func (s *FilesService) Download(ctx context.Context, fileName string) (string, error) {
	decompressedFileName, err := DecompressFiles(ctx, s.envCfg.UploadFolder, fileName)
	if err != nil {
		return "", err
	}

	return decompressedFileName, nil
}
