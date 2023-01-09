package utils

import "encoding/json"

// UnmarshalJson is a wrapper for json.Unmarshal
func UnmarshalJson(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalJson is a wrapper for json.Marshal
func MarshalJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func GetVariableValue(
	variable string,
	body map[string]interface{},
) interface{} {
	keyWithoutDollarDot := variable[2:]
	flatResult := Flatten(InterfaceToMapStringInterface(body))
	return flatResult[keyWithoutDollarDot]
}
