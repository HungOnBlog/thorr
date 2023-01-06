package utils

// Flatten flattens a nested map[string]interface{} into a single map[string]interface{}
// Example:
//
//	{
//		"foo": {
//			"bar": "baz",
//			"inside": {
//				"deep": "value"
//			}
//		}
//	}
//
// -->
//
//	{
//		"foo.bar": "baz",
//		"foo.inside.deep": "value"
//	}
func Flatten(j map[string]interface{}) map[string]interface{} {
	flat := make(map[string]interface{})
	flatten(flat, "", j)
	return flat
}

func flatten(flat map[string]interface{}, prefix string, j map[string]interface{}) {
	for k, v := range j {
		if m, ok := v.(map[string]interface{}); ok {
			flatten(flat, prefix+k+".", m)
		} else {
			flat[prefix+k] = v
		}
	}
}
