package utils

import (
	"os"
)

func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		return []byte{}
	}

	return content
}
