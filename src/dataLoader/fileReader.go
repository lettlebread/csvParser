package dataLoader

import (
	"io/ioutil"
)

type FileReader struct {
	filePath string
}

func (f *FileReader) LoadDataString() (string, error) {
	dat, err := ioutil.ReadFile(f.filePath)
	return string(dat), err
}
