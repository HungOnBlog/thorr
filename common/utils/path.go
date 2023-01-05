package utils

import "os"

// Check is the path is a file and not a directory.
// If the path is a file, check if the file is exists or not.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Return all the files in the directory.
func LoadAllFilePathsIn(path string) []string {
	os.Chdir(path)
	files, err := os.ReadDir(".")
	if err != nil {
		return []string{}
	}

	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, file.Name())
	}

	return filePaths
}
