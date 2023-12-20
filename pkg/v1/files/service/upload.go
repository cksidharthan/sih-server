package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

// Upload uploads the file to the server by generating a unique filename
func (s *FilesService) Upload(ctx context.Context, uploadedFile *multipart.FileHeader) error {
	s.log.Debugf("Uploading file '%s'", uploadedFile.Filename)
	return SaveFile(ctx, uploadedFile, generateFileName(uploadedFile, s.envCfg.UploadFolder))
}

// generateFileName generates a unique filename for the uploaded file
func generateFileName(uploadedFile *multipart.FileHeader, folderName string) string {
	fileName := fmt.Sprintf("%s/%s__++%d__++%s", folderName, uuid.New().String(), time.Now().Unix(), uploadedFile.Filename)
	return fileName
}
