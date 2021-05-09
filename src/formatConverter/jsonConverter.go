package formatConverter

import (
	"encoding/json"
)

type jsonConverter struct {
	args []string
}

func (j *jsonConverter) Exec(mapArr []map[string]string) (string, error) {
	res, err := json.Marshal(mapArr)
	return string(res), err
}
