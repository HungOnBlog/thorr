package utils

func MapAssign(m map[string]interface{}, key string, value interface{}) {
	m[key] = value
}

func MapAssignAll(m map[string]interface{}, values map[string]interface{}) map[string]interface{} {
	outputMap := make(map[string]interface{})

	if m == nil {
		return values
	}

	for k, v := range m {
		outputMap[k] = v
	}

	for k, v := range values {
		outputMap[k] = v
	}

	return outputMap
}
