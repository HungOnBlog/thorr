package utils

import (
	"encoding/json"
	"io"
)

func MapStringInterfaceToMapStringString(m map[string]interface{}) map[string]string {
	newMap := make(map[string]string)
	for k, v := range m {
		newMap[k] = v.(string)
	}
	return newMap
}

func ReadCloserToMapStringInterface(r io.ReadCloser) (map[string]interface{}, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
