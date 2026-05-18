package helper

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UploadFile(
	file *multipart.FileHeader,
	folder string,
) (string, int64, error) {

	src, err := file.Open()
	if err != nil {

		return "", 0, err
	}

	defer src.Close()

	// create folder
	dir := filepath.Join(
		"storage",
		folder,
	)

	err = os.MkdirAll(
		dir,
		os.ModePerm,
	)

	if err != nil {

		return "", 0, err
	}

	ext := filepath.Ext(
		file.Filename,
	)

	filename := fmt.Sprintf(
		"%s%s",
		uuid.NewString(),
		ext,
	)

	fullPath := filepath.Join(
		dir,
		filename,
	)

	dst, err := os.Create(fullPath)
	if err != nil {

		return "", 0, err
	}

	defer dst.Close()

	_, err = dst.ReadFrom(src)
	if err != nil {

		return "", 0, err
	}

	fileURL := fmt.Sprintf(
		"/storage/%s/%s",
		folder,
		filename,
	)

	return fileURL, file.Size, nil
}
