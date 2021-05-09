package dataLoader

import (
	"errors"
)

type DataLoader interface {
	Exec() (string, error)
}

func NewDataLoader(loaderType string, args []string) (DataLoader, error) {
	var reader DataLoader
	var err error

	switch loaderType {
	case "file":
		reader = &FileReader{args}
	default:
		err = errors.New("Invalid loader type")
	}
	return reader, err
}
