package formatConverter

import (
	"encoding/json"
)

type jsonConverter struct{}

func (j *jsonConverter) Convert(mapArr []map[string]string) (string, error) {
	res, err := json.Marshal(mapArr)
	return string(res), err
}
