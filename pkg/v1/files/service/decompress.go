package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ulikunitz/xz"
	"io"
	"os"
	"strings"
)

func DecompressFiles(ctx context.Context, uploadFolder, filePath string) (string, error) {
	// Open the input file
	inputFile, err := os.Open(fmt.Sprintf("%s/%s", uploadFolder, filePath))
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	fileNameSplit := strings.Split(filePath, ".")

	// Create the output file
	outputFile, err := os.Create(fmt.Sprintf("/tmp/%s.%s", uuid.New().String(), fileNameSplit[len(fileNameSplit)-1]))
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Create an LZMA reader
	reader, err := xz.NewReader(inputFile)
	if err != nil {
		return "", err
	}

	// Copy data from the LZMA reader to the output file
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return "", err
	}

	return outputFile.Name(), nil
}
