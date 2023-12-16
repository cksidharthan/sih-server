package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveFile saves the uploaded file to the destination as per the filename.
func SaveFile(ctx context.Context, file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(fileName), 0750); err != nil {
		return err
	}

	out, err := os.Create(fileName)
	if err != nil {

		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
