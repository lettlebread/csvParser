package formatConverter

import (
	"errors"
)

type FormatConverter interface {
	Convert(mapArr []map[string]string) (string, error)
}

func New(args map[string]string) (FormatConverter, error) {
	var err error
	var converterType string

	if val, ok := args["e"]; ok && val != "" {
		converterType = val
	} else {
		return nil, errors.New("Invalid output format type")
	}

	switch converterType {
	case "json":
		return &jsonConverter{}, err
	default:
		return nil, errors.New("Invalid output format type")
	}
}
