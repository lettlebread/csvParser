package dataLoader

import (
	"errors"
)

type DataLoader interface {
	LoadDataString() (string, error)
}

func New(args map[string]string) (DataLoader, error) {
	var err error
	var loaderType, filePath string

	if val, ok := args["f"]; ok && val != "" {
		loaderType = "file"
		filePath = val
	} else {
		return nil, errors.New("Invalid file path")
	}

	switch loaderType {
	case "file":
		return &FileReader{filePath}, err
	default:
		return nil, errors.New("Invalid loader type")
	}
}
