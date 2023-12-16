package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
)

// Upload uploads the file to the server by generating a unique filename
func (s *FilesService) Upload(ctx context.Context, uploadedFile *multipart.FileHeader) error {
	s.log.Debugf("Uploading file '%s'", uploadedFile.Filename)
	return SaveFile(ctx, uploadedFile, generateFileName(uploadedFile, s.envCfg.UploadFolder))
}

// generateFileName generates a unique filename for the uploaded file
func generateFileName(uploadedFile *multipart.FileHeader, folderName string) string {
	fileName := fmt.Sprintf("%s/%s__++%s", folderName, uuid.New().String(), uploadedFile.Filename)
	return fileName
}
