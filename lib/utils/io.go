package utils

import (
	"os"
)

func AppendLine(path string, text string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(text + "\n"); err != nil {
		return err
	}

	return nil
}
