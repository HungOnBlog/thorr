package utils

import (
	"strings"
)

// Return request path placeholder
// Example: /users/:id => [":id"]
// Example: /users/:id/:name => [":id", ":name"]
// Example: /users/:id/phone/:phone_id => [":id", ":phone_id"]
func GetPathPlaceHolders(path string) []string {
	arr := SplitBy(path, "/")
	var placeholder []string
	for _, v := range arr {
		if strings.HasPrefix(v, ":") {
			placeholder = append(placeholder, v)
		}
	}

	return placeholder
}

func CountNumberOfPathsWillBeGenerated(path string, pathParams map[string][]string) int {
	placeholders := GetPathPlaceHolders(path)
	numberOfPlaceholders := len(placeholders)
	if numberOfPlaceholders == 0 {
		return 1
	}

	numberOfPaths := 1
	for _, v := range placeholders {
		keyRemovedColon := strings.TrimPrefix(v, ":")
		numberOfPaths *= len(pathParams[keyRemovedColon])
	}

	return numberOfPaths
}

func GeneratePaths(path string, placeholders []string, pathParams map[string][]string) []string {

	var paths []string

	if len(placeholders) == 0 {
		paths = append(paths, path)
		return paths
	}

	for i, pl := range placeholders {
		if i == 0 {
			paths = GeneratePathForSinglePlaceholder([]string{path}, pl, pathParams)
		} else {
			paths = GeneratePathForSinglePlaceholder(paths, pl, pathParams)
		}
	}

	return paths
}

func GeneratePathForSinglePlaceholder(paths []string, placeholder string, pathParams map[string][]string) []string {
	var genPaths []string

	keyRemovedColon := strings.TrimPrefix(placeholder, ":")
	values := pathParams[keyRemovedColon]
	for _, path := range paths {
		for _, value := range values {
			genPaths = append(genPaths, strings.Replace(path, placeholder, value, 1))
		}
	}

	return genPaths
}
