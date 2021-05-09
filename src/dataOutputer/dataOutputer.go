package dataOutputer

import (
	"errors"
)

type DataOutputer interface {
	OutputData(data string) error
}

func NewDataOutputer(args map[string]string) (DataOutputer, error) {
	var outputer DataOutputer
	var err error
	var outputerType, outputPath string

	if val, ok := args["o"]; ok {
		outputerType = "file"
		outputPath = val
	}

	switch outputerType {
	case "file":
		outputer = &fileWriter{outputPath}
	default:
		err = errors.New("Invalid output data type")
	}
	return outputer, err
}
