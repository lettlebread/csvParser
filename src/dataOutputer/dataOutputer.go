package dataOutputer

import (
	"errors"
)

type DataOutputer interface {
	Exec(data string) error
}

func NewDataOutputer(outputType string, args []string) (DataOutputer, error) {
	var outputer DataOutputer
	var err error

	switch outputType {
	case "file":
		outputer = &fileWriter{args}
	default:
		err = errors.New("Invalid output data type")
	}
	return outputer, err
}
