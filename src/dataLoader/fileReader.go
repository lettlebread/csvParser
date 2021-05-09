package dataLoader

import (
	"errors"
	"io/ioutil"
)

type FileReader struct {
	args []string
}

func (f *FileReader) Exec() (string, error) {
	filePath := f.args[0]

	if len(filePath) == 0 {
		return "", errors.New("Invalid fie path")
	}

	dat, err := ioutil.ReadFile(filePath)
	return string(dat), err
}
