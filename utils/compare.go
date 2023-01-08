package utils

import "fmt"

func CompareMap(expected map[string]interface{}, actual map[string]interface{}) error {
	for k, v := range expected {
		if actual[k] != v {
			return fmt.Errorf("expected::%v, got::%v", v, actual[k])
		}
	}

	return nil
}
