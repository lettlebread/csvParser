package formatConverter

import (
	"errors"
)

type FormatConverter interface {
	Exec(mapArr []map[string]string) (string, error)
}

func NewFormatConverter(converterType string, args []string) (FormatConverter, error) {
	var converter FormatConverter
	var err error

	switch converterType {
	case "json":
		converter = &jsonConverter{args}
	default:
		err = errors.New("Invalid output format type")
	}
	return converter, err
}
