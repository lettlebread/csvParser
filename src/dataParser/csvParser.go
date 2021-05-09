package dataParser

import (
	"errors"
)

type csvParser struct{}

func (c *csvParser) ParseToMapArray(data string) ([]map[string]string, error) {
	lineParser := newLineParser("\n")
	fieldParser := newFieldParser(",")
	res := []map[string]string{}

	lines := lineParser(data)
	fieldName := fieldParser(lines[0])

	for _, line := range lines[1:] {
		fields := fieldParser(line)

		newMap, err := arrayToMap(fieldName, fields)

		if err != nil {
			return nil, err
		}

		res = append(res, newMap)
	}

	return res, nil
}

func arrayToMap(fieldName []string, fields []string) (map[string]string, error) {
	newMap := map[string]string{}

	if len(fields) != len(fieldName) {
		return nil, errors.New("Invalid data format")
	}

	for i, val := range fieldName {
		newMap[val] = fields[i]
	}

	return newMap, nil
}
