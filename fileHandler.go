package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func BuildFileDictionary(inputDir string) (map[string]string, error) {
	fileDictionary := make(map[string]string)
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".qtml") {
			componentName := strings.TrimSuffix(info.Name(), ".qtml")
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if _, exists := fileDictionary[componentName]; exists {
				return errors.New("Found two components with the same name: " + componentName)
			}
			stringContent := convertToDataAttributes(string(content))
			fileDictionary[componentName] = stringContent
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return fileDictionary, nil
}
