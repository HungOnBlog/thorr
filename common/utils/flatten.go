package utils

// FlattenJson flattens a json object
// Example:
//
//	{
//		"key1": "value1",
//		"key2": {
//			"key3": "value3"
//		}
//	}
//	->
//	{
//		"key1": "value1",
//		"key2.key3": "value3"
//	}
func FlattenJson(j map[string]interface{}) map[string]interface{} {
	flatObject := make(map[string]interface{})
	for k, v := range j {
		flatten(flatObject, k, v)
	}
	return flatObject
}

func flatten(flatObject map[string]interface{}, key string, value interface{}) {
	switch v := value.(type) {
	case map[string]interface{}:
		for k, v := range v {
			flatten(flatObject, key+"."+k, v)
		}
	default:
		flatObject[key] = v
	}
}
