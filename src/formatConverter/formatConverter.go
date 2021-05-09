package formatConverter

import (
	"errors"
)

type FormatConverter interface {
	Convert(mapArr []map[string]string) (string, error)
}

func NewFormatConverter(args map[string]string) (FormatConverter, error) {
	var converter FormatConverter
	var err error
	var converterType string

	if val, ok := args["e"]; ok {
		converterType = val
	}

	switch converterType {
	case "json":
		converter = &jsonConverter{}
	default:
	}

	if converter == nil {
		err = errors.New("Invalid output format type")
	}

	return converter, err
}
