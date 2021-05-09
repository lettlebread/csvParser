package dataParser

import (
	"fmt"
)

type csvParser struct{}

func (c *csvParser) ParseToMapArray(data string) ([]map[string]string, error) {
	lineParser := setLineParser("\n")
	fieldParser := setFieldParser(",")
	res := make([]map[string]string, 0)

	lines := lineParser(data)
	fieldName := fieldParser(lines[0])
	fieldNum := len(fieldName)

	for _, line := range lines[1:] {
		newMap := make(map[string]string)
		fields := fieldParser(line)

		if len(fields) != fieldNum {
			return nil, fmt.Errorf("invalid data format")
		}

		for i, val := range fieldName {
			newMap[val] = fields[i]
		}
		res = append(res, newMap)
	}

	return res, nil
}
