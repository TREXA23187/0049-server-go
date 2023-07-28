package file_service

import (
	"io"
	"net/http"
)

func GetQiNiuFileBytes(filePath string) ([]byte, error) {

	// Send GET request to OSS URL
	resp, err := http.Get(filePath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
