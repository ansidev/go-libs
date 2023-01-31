package io

import (
	"os"
)

func CreateDirIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirPath, 0755)
		if errDir != nil {
			return err
		}
	}

	return nil
}

func TruncateFile(filePath string) error {
	err := os.Truncate(filePath, 0)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(filePath string, content []byte) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return err
	}

	return nil
}
