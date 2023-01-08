package validators

import (
	"fmt"
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

func (t *TypeValidator) Validate(expected interface{}, actual interface{}) error {
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
		return fmt.Errorf("type::%s is not supported", expectedType)
	}
}

func (t *TypeValidator) validateString(command string, value string, actual interface{}) error {
	if !utils.IsString(actual) {
		return fmt.Errorf("expected::string, got::%T", actual)
	}

	if command == "" {
		return nil
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
		return fmt.Errorf("command::%s is not supported", command)
	}
}

// Validate email
// Example: email::true
func (t *TypeValidator) validateEmail(actual interface{}) error {
	emailRegex := `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$`
	err := t.validateRegex(emailRegex, actual)
	if err != nil {
		return fmt.Errorf("expected::string::email, got::%v", actual)
	}

	return nil
}

// Validate length
// Example: length::5-10
func (t *TypeValidator) validateLength(value string, actual interface{}) error {
	length := len(utils.InterfaceToString(actual))
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToInt(minMax[0])
	max := utils.StringToInt(minMax[1])

	isInRange := length >= min && length <= max
	if !isInRange {
		return fmt.Errorf("expected::string::length::%v, got::%v", value, length)
	}

	return nil
}

// Validate UUID
// Example of UUID: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
func (t *TypeValidator) validateUUID(actual interface{}) error {
	uuidRegex := `^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	err := t.validateRegex(uuidRegex, actual)
	if err != nil {
		return fmt.Errorf("expected::string::uuid, got::%v", actual)
	}

	return nil
}

// Validate URL
// Example of URL: https://www.google.com
func (t *TypeValidator) validateURL(actual interface{}) error {
	urlRegex := `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	err := t.validateRegex(urlRegex, actual)
	if err != nil {
		return fmt.Errorf("expected::string::url, got::%v", actual)
	}

	return nil
}

// Validate base64
// Example of base64: aGVsbG8gd29ybGQ=
func (t *TypeValidator) validateBase64(actual interface{}) error {
	base64Regex := `^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`
	err := t.validateRegex(base64Regex, actual)
	if err != nil {
		return fmt.Errorf("expected::string::base64, got::%v", actual)
	}

	return nil
}

// Validate regex
// Example: regex::^[a-z]+$
func (t *TypeValidator) validateRegex(regex string, actual interface{}) error {
	isMatch := utils.MatchRegex(regex, utils.InterfaceToString(actual))
	if !isMatch {
		return fmt.Errorf("expected::string::regex::%v, got::%v", regex, actual)
	}

	return nil
}

func (t *TypeValidator) validateInt(command string, value string, actual interface{}) error {
	if !utils.IsInt(actual) {
		return fmt.Errorf("expected::int, got::%T", actual)
	}

	if command == "" {
		return nil
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
		return fmt.Errorf("command is not supported")
	}
}

// Validate int range
// Example: range::5-10
func (t *TypeValidator) validateIntRange(value string, actual interface{}) error {
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToInt(minMax[0])
	max := utils.StringToInt(minMax[1])

	isInRange := actual.(int) >= min && actual.(int) <= max
	if !isInRange {
		return fmt.Errorf("expected::int::range::%v, got::%v", value, actual)
	}

	return nil
}

// Validate int logic
// Example: lt::5
func (t *TypeValidator) validateIntLogic(command string, value string, actual interface{}) error {
	compare := utils.StringToInt(value)
	var isLogicPassed bool

	switch command {
	case "lt":
		isLogicPassed = actual.(int) < compare
	case "lte":
		isLogicPassed = actual.(int) <= compare
	case "gt":
		isLogicPassed = actual.(int) > compare
	case "gte":
		isLogicPassed = actual.(int) >= compare
	default:
		return fmt.Errorf("command is not supported")
	}

	if !isLogicPassed {
		return fmt.Errorf("expected::int::%v::%v, got::%v", command, value, actual)
	}

	return nil
}

func (t *TypeValidator) validateFloat(command string, value string, actual interface{}) error {
	if !utils.IsFloat(actual) {
		return fmt.Errorf("expected::float, got::%T", actual)
	}

	if command == "" {
		return nil
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
		return fmt.Errorf("command is not supported")
	}
}

// Validate float range
// Example: range::5.5-10.5
func (t *TypeValidator) validateFloatRange(value string, actual interface{}) error {
	minMax := utils.SplitBy(value, "-")
	min := utils.StringToFloat(minMax[0])
	max := utils.StringToFloat(minMax[1])

	isInRange := actual.(float64) >= min && actual.(float64) <= max
	if !isInRange {
		return fmt.Errorf("expected::float::range::%v, got::%v", value, actual)
	}

	return nil
}

// Validate float logic
// Example: lt::5.5
func (t *TypeValidator) validateFloatLogic(command string, value string, actual interface{}) error {
	compare := utils.StringToFloat(value)
	var isLogicPassed bool

	switch command {
	case "lt":
		isLogicPassed = actual.(float64) < compare
	case "lte":
		isLogicPassed = actual.(float64) <= compare
	case "gt":
		isLogicPassed = actual.(float64) > compare
	case "gte":
		isLogicPassed = actual.(float64) >= compare
	default:
		return fmt.Errorf("command is not supported")
	}

	if !isLogicPassed {
		return fmt.Errorf("expected::float::%v::%v, got::%v", command, value, actual)
	}

	return nil
}

func (t *TypeValidator) validateBool(command string, value string, actual interface{}) error {
	if !utils.IsBool(actual) {
		return fmt.Errorf("expected::bool, got::%T", actual)
	}

	return nil
}

func (t *TypeValidator) validateDate(command string, value string, actual interface{}) error {
	if command == "" {
		_, err := time.Parse("2006-01-02", utils.InterfaceToString(actual))
		if err != nil {
			return fmt.Errorf("expected::date, got::%T default format is (yyyy-MM-dd)", actual)
		}

		return nil
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
		return fmt.Errorf("command is not supported")
	}
}

// Validate date before
// Example: before::2018-01-01
func (t *TypeValidator) validateDateBefore(value string, actual interface{}) error {
	date, err := utils.StringToDate(value)
	if err != nil {
		return err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	isBefore := actualDate.Before(date)
	if !isBefore {
		return fmt.Errorf("expected::date::before::%v, got::%v", value, actual)
	}

	return nil
}

// Validate date after
// Example: after::2018-01-01
func (t *TypeValidator) validateDateAfter(value string, actual interface{}) error {
	date, err := utils.StringToDate(value)
	if err != nil {
		return err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	isAfter := actualDate.After(date)
	if !isAfter {
		return fmt.Errorf("expected::date::after::%v, got::%v", value, actual)
	}

	return nil
}

// Validate date range
// Example: range::2018-01-01 to 2018-01-31
func (t *TypeValidator) validateDateRange(value string, actual interface{}) error {
	dates := utils.SplitBy(value, "to")
	from, err := utils.StringToDate(dates[0])
	if err != nil {
		return err
	}

	to, err := utils.StringToDate(dates[1])
	if err != nil {
		return err
	}

	actualDate, err := utils.StringToDate(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	inRange := actualDate.After(from) && actualDate.Before(to)
	onFrom := actualDate.Equal(from)
	onTo := actualDate.Equal(to)

	isOnRange := inRange || onFrom || onTo
	if !isOnRange {
		return fmt.Errorf("expected::date::range::%v, got::%v", value, actual)
	}

	return nil
}

// Validate date format
// Example: format::2023/01/31
func (t *TypeValidator) validateDateFormat(format string, actual interface{}) error {
	_, err := utils.StringToDateWithFormat(utils.InterfaceToString(actual), format)
	if err != nil {
		return fmt.Errorf("expected::date::format::%v, got::%v", format, actual)
	}

	return nil
}

// Validate date utc
// Example UTC Date Format: 2018-01-01T00:00:00Z
// Example UTC Date Format: 2018-01-01T00:00:00+00:00
// Example UTC Date Format: 2018-01-01T00:00:00+0000
func (t *TypeValidator) validateDateUTC(actual interface{}) error {
	_, err := utils.StringToUTC(utils.InterfaceToString(actual))
	if err != nil {
		return fmt.Errorf("expected::date::utc, got::%v", actual)
	}

	return nil
}

// Validate date ISO8601
// Example ISO8601 Date Format: 2018-01-01T00:00:00+00:00
func (t *TypeValidator) validateDateISO8601(actual interface{}) error {
	_, err := utils.StringIso8601ToTime(utils.InterfaceToString(actual))
	if err != nil {
		return fmt.Errorf("expected::date::iso8601, got::%v", actual)
	}

	return nil
}

func (t *TypeValidator) validateTime(command string, value string, actual interface{}) error {
	if !utils.IsTimeString(actual) {
		return fmt.Errorf("expected::time, got::%T default format is HH:mm:ss", actual)
	}

	if command == "" {
		return nil
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
		return fmt.Errorf("command is not supported")
	}
}

// Validate time before
// Example: before::12:00:00
func (t *TypeValidator) validateTimeBefore(value string, actual interface{}) error {
	time, err := utils.StringToTime(value)
	if err != nil {
		return err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	isBefore := actualTime.Before(time)
	if !isBefore {
		return fmt.Errorf("expected::time::before::%v, got::%v", value, actual)
	}

	return nil
}

// Validate time after
// Example: after::12:00:00
func (t *TypeValidator) validateTimeAfter(value string, actual interface{}) error {
	time, err := utils.StringToTime(value)
	if err != nil {
		return err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	isAfter := actualTime.After(time)
	if !isAfter {
		return fmt.Errorf("expected::time::after::%v, got::%v", value, actual)
	}

	return nil
}

// Validate time range
// Example: range::12:00:00 to 13:00:00
func (t *TypeValidator) validateTimeRange(value string, actual interface{}) error {
	times := utils.SplitBy(value, "to")
	from, err := utils.StringToTime(times[0])
	if err != nil {
		return err
	}

	to, err := utils.StringToTime(times[1])
	if err != nil {
		return err
	}

	actualTime, err := utils.StringToTime(utils.InterfaceToString(actual))
	if err != nil {
		return err
	}

	isInRange := actualTime.After(from) && actualTime.Before(to)
	onFrom := actualTime.Equal(from)
	onTo := actualTime.Equal(to)

	isOnRange := isInRange || onFrom || onTo
	if !isOnRange {
		return fmt.Errorf("expected::time::range::%v, got::%v", value, actual)
	}

	return nil
}

// Validate time format
// Example: format::15:04:05
func (t *TypeValidator) validateTimeFormat(format string, actual interface{}) error {
	_, err := utils.StringToTimeWithFormat(utils.InterfaceToString(actual), format)
	if err != nil {
		return fmt.Errorf("expected::time::format::%v, got::%v", format, actual)
	}

	return nil
}
