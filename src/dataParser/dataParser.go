package dataParser

import (
	"errors"
	"strings"
)

type DataParser interface {
	ParseToMapArray(src string) ([]map[string]string, error)
}

func NewDataParser(args map[string]string) (DataParser, error) {
	var parser DataParser
	var err error
	var parserType string

	if val, ok := args["m"]; ok {
		parserType = val
	}

	switch parserType {
	case "csv":
		parser = &csvParser{}
	default:
	}

	if parser == nil {
		err = errors.New("Invalid parser type")
	}

	return parser, err
}

func setLineParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		return strings.Split(line, delimiter)
	}
}

func setFieldParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		line = strings.TrimSuffix(line, "\n")
		return strings.Split(line, delimiter)
	}
}
