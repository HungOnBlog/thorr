package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
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

func InterfaceToString(value interface{}) string {
	return value.(string)
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return intValue
}

func StringToFloat(value string) float64 {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return floatValue
}

func StringToDate(value string) (time.Time, error) {
	dateValue, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringToDateWithFormat(value string, format string) (time.Time, error) {
	dateValue, err := time.Parse(format, value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringToUTC(value string) (time.Time, error) {
	dateValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringToTime(value string) (time.Time, error) {
	dateValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringToTimeWithFormat(value string, format string) (time.Time, error) {
	dateValue, err := time.Parse(format, value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringIso8601ToTime(value string) (time.Time, error) {
	dateValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}

	return dateValue, nil
}

func StringUnixToTime(value string) (time.Time, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(intValue, 0), nil
}

func HttpHeaderToMapStringString(header http.Header) map[string]string {
	newMap := make(map[string]string)
	for k, v := range header {
		newMap[k] = v[0]
	}
	return newMap
}

func MapStringStringToMapStringInterface(m map[string]string) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func InterfaceToFloat64(value interface{}) float64 {
	return value.(float64)
}
