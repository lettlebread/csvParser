package dataLoader

import (
	"errors"
	"io/ioutil"
)

type FileReader struct {
	filePath string
}

func (f *FileReader) LoadDataString() (string, error) {
	if len(f.filePath) == 0 {
		return "", errors.New("Invalid file path")
	}

	dat, err := ioutil.ReadFile(f.filePath)
	return string(dat), err
}
