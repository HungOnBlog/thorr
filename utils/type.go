package utils

func IsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func IsInt(value interface{}) bool {
	_, ok := value.(int)
	return ok
}

func IsFloat(value interface{}) bool {
	_, ok := value.(float64)
	return ok
}

func IsBool(value interface{}) bool {
	_, ok := value.(bool)
	return ok
}

func IsArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func IsMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

func IsDateString(value interface{}) bool {
	return IsString(value) && MatchRegex(`^\d{4}-\d{2}-\d{2}$`, value.(string))
}

func IsDateTimeString(value interface{}) bool {
	return IsString(value) && MatchRegex(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`, value.(string))
}

func IsTimeString(value interface{}) bool {
	return IsString(value) && MatchRegex(`^\d{2}:\d{2}:\d{2}$`, value.(string))
}
