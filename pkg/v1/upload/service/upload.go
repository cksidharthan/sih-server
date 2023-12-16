package service

import "mime/multipart"

func (s *UploadService) Upload(uploadedFile *multipart.FileHeader) error {
	return SaveFile(uploadedFile, "temp/"+uploadedFile.Filename)
}
