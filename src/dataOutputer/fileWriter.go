package dataOutputer

import (
	"os"
)

type fileWriter struct {
	outputPath string
}

func (f *fileWriter) OutputData(data string) error {
	newFile, err := os.Create(f.outputPath)
	defer newFile.Close()

	if err != nil {
		return err
	}

	_, err = newFile.WriteString(data)

	if err != nil {
		return err
	}

	newFile.Sync()

	return err
}
