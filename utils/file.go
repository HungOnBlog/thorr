package utils

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ListAllFilesPathInDir(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func IsJsonFile(filePath string) bool {
	return filePath[len(filePath)-5:] == ".json"
}

func IsYamlFile(filePath string) bool {
	yaml := filePath[len(filePath)-5:] == ".yaml"
	yml := filePath[len(filePath)-4:] == ".yml"
	return yaml || yml
}

func IsFile(filePath string) bool {
	file, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return !file.IsDir()
}
