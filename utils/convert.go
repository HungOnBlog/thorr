package utils

func MapStringInterfaceToMapStringString(m map[string]interface{}) map[string]string {
	newMap := make(map[string]string)
	for k, v := range m {
		newMap[k] = v.(string)
	}
	return newMap
}
