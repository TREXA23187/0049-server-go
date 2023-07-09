package ctype

import "encoding/json"

type FileType string

const (
	DataFileType  FileType = "data"
	ModelFileType FileType = "model"
)

func (f FileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f)
}
