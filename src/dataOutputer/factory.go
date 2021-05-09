package dataOutputer

import (
	"errors"
)

type DataOutputer interface {
	OutputData(data string) error
}

func New(args map[string]string) (DataOutputer, error) {
	var err error
	var outputerType, outputPath string

	if val, ok := args["o"]; ok && val != "" {
		outputerType = "file"
		outputPath = val
	} else {
		return nil, errors.New("Invalid output data path")
	}

	switch outputerType {
	case "file":
		return &fileWriter{outputPath}, err
	default:
		return nil, errors.New("Invalid output data type")
	}
}
