package utils

import "io/ioutil"

func ReadJsonFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}
	}

	return content
}
