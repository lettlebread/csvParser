package dataLoader

import (
	"errors"
)

type DataLoader interface {
	LoadDataString() (string, error)
}

func NewDataLoader(args map[string]string) (DataLoader, error) {
	var reader DataLoader
	var err error
	var loaderType, filePath string

	if val, ok := args["f"]; ok {
		loaderType = "file"
		filePath = val
	}

	switch loaderType {
	case "file":
		reader = &FileReader{filePath}
	default:
	}

	if reader == nil {
		err = errors.New("Invalid loader type")
	}

	return reader, err
}
