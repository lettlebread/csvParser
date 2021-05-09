package dataParser

import (
	"errors"
	"strings"
)

type DataParser interface {
	ParseToMapArray(src string) ([]map[string]string, error)
}

func New(args map[string]string) (DataParser, error) {
	var err error
	var parserType string

	if val, ok := args["m"]; ok && val != "" {
		parserType = val
	} else {
		return nil, errors.New("Invalid parser type")
	}

	switch parserType {
	case "csv":
		return &csvParser{}, err
	default:
		return nil, errors.New("Invalid parser type")
	}
}

func newLineParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		return strings.Split(line, delimiter)
	}
}

func newFieldParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		line = strings.TrimSuffix(line, "\n")
		return strings.Split(line, delimiter)
	}
}
