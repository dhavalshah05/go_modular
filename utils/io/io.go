package io

import "os"

func ReadFile(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}

func WriteFile(filePath string, data string) error {
	return os.WriteFile(filePath, []byte(data), 0644)
}
