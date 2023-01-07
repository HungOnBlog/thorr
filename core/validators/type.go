package validators

import (
	"errors"
	"time"

	"github.com/HungOnBlog/thorr/utils"
)

type TypeValidator struct{}

func NewTypeValidator() *TypeValidator {
	return &TypeValidator{}
}

// Get all definitions from expected string
// Expected string format: type::command::value
// Example: string::min::5
func (t *TypeValidator) parseExpected(expected string) (string, string, string) {
	def := utils.SplitBy(expected, "::")
	typeDef := def[0]
	var command string
	var value string
	if len(def) > 1 {
		command = def[1]
	}

	if len(def) > 2 {
		value = def[2]
	}

	return typeDef, command, value
}

func (t *TypeValidator) Validate(expected interface{}, actual interface{}) (bool, error) {
	expectedType, command, value := t.parseExpected(utils.InterfaceToString(expected))

	switch expectedType {
	case "string":
		return t.validateString(command, value, actual)
	case "int":
		return t.validateInt(command, value, actual)
	case "float":
		return t.validateFloat(command, value, actual)
	case "bool":
		return t.validateBool(command, value, actual)
	case "date":
		return t.validateDate(command, value, actual)
	case "time":
		return t.validateTime(command, value, actual)
	default:
		return false, errors.New("command is not supported")
	}
}

func (t *TypeValidator) validateString(command string, value string, actual interface{}) (bool, error) {
	if !utils.IsString(actual) {
		return false, nil
	}

	if command == "" {
		return true, nil
	}

	switch command {
	case "email":
		return t.validateEmail(actual)
	case "length":
		return t.validateLength(value, actual)
	case "uuid":
		return t.validateUUID(actual)
	case "url":
		return t.validateURL(actual)
	case "base64":
		return t.validateBase64(actual)
	case "regex":
		return t.validateRegex(value, actual)
	default:
		return false, errors.New("command is not supported")
	}
}

// Validate email
// Example: email::true
func (t *TypeValidator) validateEmail(actual interface{}) (bool, error) {
	emailRegex := `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$`
	return t.validateRegex(emailRegex, actual)
}

// Validate length
// Example: length::5-10
func (t *TypeValidator) validateLength(value string, actual interface{}) (bool, error) {
	length := len(utils.InterfaceToString(actual))
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToInt(minMax[0])
	max := utils.StringToInt(minMax[1])

	return length >= min && length <= max, nil
}

// Validate UUID
// Example of UUID: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
func (t *TypeValidator) validateUUID(actual interface{}) (bool, error) {
	uuidRegex := `^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	return t.validateRegex(uuidRegex, actual)
}

// Validate URL
// Example of URL: https://www.google.com
func (t *TypeValidator) validateURL(actual interface{}) (bool, error) {
	urlRegex := `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	return t.validateRegex(urlRegex, actual)
}

// Validate base64
// Example of base64: aGVsbG8gd29ybGQ=
func (t *TypeValidator) validateBase64(actual interface{}) (bool, error) {
	base64Regex := `^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`
	return t.validateRegex(base64Regex, actual)
}

// Validate regex
// Example: regex::^[a-z]+$
func (t *TypeValidator) validateRegex(regex string, actual interface{}) (bool, error) {
	return utils.MatchRegex(regex, utils.InterfaceToString(actual)), nil
}

func (t *TypeValidator) validateInt(command string, value string, actual interface{}) (bool, error) {
	if !utils.IsInt(actual) {
		return false, nil
	}

	if command == "" {
		return true, nil
	}

	switch command {
	case "range":
		return t.validateIntRange(value, actual)
	case "lt":
		return t.validateIntLogic("lt", value, actual)
	case "lte":
		return t.validateIntLogic("lte", value, actual)
	case "gt":
		return t.validateIntLogic("gt", value, actual)
	case "gte":
		return t.validateIntLogic("gte", value, actual)
	default:
		return false, errors.New("command is not supported")
	}
}

// Validate int range
// Example: range::5-10
func (t *TypeValidator) validateIntRange(value string, actual interface{}) (bool, error) {
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToInt(minMax[0])
	max := utils.StringToInt(minMax[1])

	return actual.(int) >= min && actual.(int) <= max, nil
}

// Validate int logic
// Example: lt::5
func (t *TypeValidator) validateIntLogic(command string, value string, actual interface{}) (bool, error) {
	compare := utils.StringToInt(value)

	switch command {
	case "lt":
		return actual.(int) < compare, nil
	case "lte":
		return actual.(int) <= compare, nil
	case "gt":
		return actual.(int) > compare, nil
	case "gte":
		return actual.(int) >= compare, nil
	default:
		return false, errors.New("command is not supported")
	}
}

func (t *TypeValidator) validateFloat(command string, value string, actual interface{}) (bool, error) {
	if !utils.IsFloat(actual) {
		return false, nil
	}

	if command == "" {
		return true, nil
	}

	switch command {
	case "range":
		return t.validateFloatRange(value, actual)
	case "lt":
		return t.validateFloatLogic("lt", value, actual)
	case "lte":
		return t.validateFloatLogic("lte", value, actual)
	case "gt":
		return t.validateFloatLogic("gt", value, actual)
	case "gte":
		return t.validateFloatLogic("gte", value, actual)
	default:
		return false, errors.New("command is not supported")
	}
}

// Validate float range
// Example: range::5.5-10.5
func (t *TypeValidator) validateFloatRange(value string, actual interface{}) (bool, error) {
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToFloat(minMax[0])
	max := utils.StringToFloat(minMax[1])

	return actual.(float64) >= min && actual.(float64) <= max, nil
}

// Validate float logic
// Example: lt::5.5
func (t *TypeValidator) validateFloatLogic(command string, value string, actual interface{}) (bool, error) {
	compare := utils.StringToFloat(value)

	switch command {
	case "lt":
		return actual.(float64) < compare, nil
	case "lte":
		return actual.(float64) <= compare, nil
	case "gt":
		return actual.(float64) > compare, nil
	case "gte":
		return actual.(float64) >= compare, nil
	default:
		return false, errors.New("command is not supported")
	}
}

func (t *TypeValidator) validateBool(command string, value string, actual interface{}) (bool, error) {
	if !utils.IsBool(actual) {
		return false, nil
	}

	return true, nil
}

func (t *TypeValidator) validateDate(command string, value string, actual interface{}) (bool, error) {
	if command == "" {
		_, err := time.Parse("2006-01-02", utils.InterfaceToString(actual))
		if err != nil {
			return false, errors.New("invalid date format (YYYY-MM-DD)")
		}

		return true, nil
	}

	switch command {
	case "before":
		return t.validateDateBefore(value, actual)
	case "after":
		return t.validateDateAfter(value, actual)
	case "range":
		return t.validateDateRange(value, actual)
	case "format":
		return t.validateDateFormat(value, actual)
	case "utc":
		return t.validateDateUTC(actual)
	case "iso8601":
		return t.validateDateISO8601(actual)
	default:
		return false, errors.New("command is not supported")
	}
}

// Validate date before
// Example: before::2018-01-01
func (t *TypeValidator) validateDateBefore(value string, actual interface{}) (bool, error) {
	date, err := utils.StringToDate(value)
	if err != nil {
		return false, err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}
	return actualDate.Before(date), nil
}

// Validate date after
// Example: after::2018-01-01
func (t *TypeValidator) validateDateAfter(value string, actual interface{}) (bool, error) {
	date, err := utils.StringToDate(value)
	if err != nil {
		return false, err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}
	return actualDate.After(date), nil
}

// Validate date range
// Example: range::2018-01-01 to 2018-01-31
func (t *TypeValidator) validateDateRange(value string, actual interface{}) (bool, error) {
	dates := utils.SplitBy(value, "to")
	from, err := utils.StringToDate(dates[0])
	if err != nil {
		return false, err
	}

	to, err := utils.StringToDate(dates[1])
	if err != nil {
		return false, err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	inRange := actualDate.After(from) && actualDate.Before(to)
	onFrom := actualDate.Equal(from)
	onTo := actualDate.Equal(to)

	return inRange || onFrom || onTo, nil
}

// Validate date format
// Example: format::2023/01/31
func (t *TypeValidator) validateDateFormat(format string, actual interface{}) (bool, error) {
	_, err := utils.StringToDateWithFormat(utils.InterfaceToString(actual), format)
	return err == nil, err
}

// Validate date utc
// Example UTC Date Format: 2018-01-01T00:00:00Z
// Example UTC Date Format: 2018-01-01T00:00:00+00:00
// Example UTC Date Format: 2018-01-01T00:00:00+0000
func (t *TypeValidator) validateDateUTC(actual interface{}) (bool, error) {
	_, err := utils.StringToUTC(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	return err == nil, err
}

// Validate date ISO8601
// Example ISO8601 Date Format: 2018-01-01T00:00:00+00:00
func (t *TypeValidator) validateDateISO8601(actual interface{}) (bool, error) {
	_, err := utils.StringIso8601ToTime(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	return err == nil, err
}

func (t *TypeValidator) validateTime(command string, value string, actual interface{}) (bool, error) {
	if !utils.IsTimeString(actual) {
		return false, nil
	}

	if command == "" {
		return true, nil
	}

	switch command {
	case "before":
		return t.validateTimeBefore(value, actual)
	case "after":
		return t.validateTimeAfter(value, actual)
	case "range":
		return t.validateTimeRange(value, actual)
	case "format":
		return t.validateTimeFormat(value, actual)
	default:
		return false, errors.New("command is not supported")
	}
}

// Validate time before
// Example: before::12:00:00
func (t *TypeValidator) validateTimeBefore(value string, actual interface{}) (bool, error) {
	time, err := utils.StringToTime(value)
	if err != nil {
		return false, err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	return actualTime.Before(time), nil
}

// Validate time after
// Example: after::12:00:00
func (t *TypeValidator) validateTimeAfter(value string, actual interface{}) (bool, error) {
	time, err := utils.StringToTime(value)
	if err != nil {
		return false, err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	return actualTime.After(time), nil
}

// Validate time range
// Example: range::12:00:00 to 13:00:00
func (t *TypeValidator) validateTimeRange(value string, actual interface{}) (bool, error) {
	times := utils.SplitBy(value, "to")
	from, err := utils.StringToTime(times[0])
	if err != nil {
		return false, err
	}

	to, err := utils.StringToTime(times[1])
	if err != nil {
		return false, err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return false, err
	}

	return actualTime.After(from) && actualTime.Before(to), nil
}

// Validate time format
// Example: format::15:04:05
func (t *TypeValidator) validateTimeFormat(format string, actual interface{}) (bool, error) {
	_, err := utils.StringToTimeWithFormat(utils.InterfaceToString(actual), format)
	return err == nil, err
}
