package utils

import (
	"os"
)

func CreateFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}
