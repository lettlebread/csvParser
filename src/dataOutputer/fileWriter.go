package dataOutputer

import (
	"os"
)

type fileWriter struct {
	args []string
}

func (f *fileWriter) Exec(data string) error {
	newFile, err := os.Create(f.args[0])

	if err != nil {
		return err
	}

	_, err = newFile.WriteString(data)

	if err != nil {
		return err
	}

	newFile.Sync()
	defer newFile.Close()

	return err
}
