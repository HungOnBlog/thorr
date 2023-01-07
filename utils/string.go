package utils

import "strings"

func SplitBy(value string, separator string) []string {
	arr := strings.Split(value, separator)
	for i, v := range arr {
		arr[i] = strings.TrimSpace(v)
	}

	return arr
}
