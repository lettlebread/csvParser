package dataParser

import (
	"errors"
	"strings"
)

type DataParser interface {
	ParseToMapArray(src string) ([]map[string]string, error)
}

func NewDataParser(parserType string, args []string) (DataParser, error) {
	var parser DataParser
	var err error

	switch parserType {
	case "csv":
		parser = &csvParser{args}
	default:
		err = errors.New("Invalid parser type")
	}
	return parser, err
}

func setLineParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		//line = strings.TrimSuffix(line, "\n")
		return strings.Split(line, delimiter)
	}
}

func setFieldParser(delimiter string) func(line string) []string {
	return func(line string) []string {
		line = strings.TrimSuffix(line, "\n")
		return strings.Split(line, delimiter)
	}
}
