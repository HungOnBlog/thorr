package models

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/validators"
	"github.com/HungOnBlog/thorr/utils"
)

type Assertion struct {
	On       string                 `json:"on" yaml:"on"`
	Check    string                 `json:"check" yaml:"check"`
	Expected interface{}            `json:"expected" yaml:"expected"`
	At       map[string]interface{} `json:"@" yaml:"@"`
}

func (a *Assertion) CheckAssertion(result Result) error {
	switch a.On {
	case "status_code":
		return a.checkStatus(result)
	case "body":
		return a.checkBody(result)
	case "header":
		return a.checkHeader(result)
	case "body::array":
		return a.checkBodyArray(result)
	default:
		return fmt.Errorf("on::%s is no supported", a.On)
	}
}

func (a *Assertion) checkBodyArray(result Result) error {
	switch a.Check {
	case "type":
		return a.checkBodyArrayType(result.Body)
	case "exist":
		return a.checkBodyArrayExist(result.Body)
	default:
		return fmt.Errorf("on::body::array command::%s is no supported", a.Check)
	}
}

func (a *Assertion) checkBodyArrayType(body interface{}) error {
	expectedBodyType := a.Expected.(map[string]interface{})
	expectedFlattenBody := utils.Flatten(expectedBodyType)
	bodyArray := body.([]interface{})
	var typeValidator validators.IValidator = validators.NewTypeValidator()
	var errs []error

	for _, body := range bodyArray {
		bodyMap := body.(map[string]interface{})
		flattenBody := utils.Flatten(bodyMap)
		for key, expectedType := range expectedFlattenBody {
			if _, ok := flattenBody[key]; !ok {
				errs = append(errs, fmt.Errorf("on::body::array key::%s is not exist", key))
				continue
			}
			err := typeValidator.Validate(expectedType, flattenBody[key])
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("on::body command::type %s", errs)
	}

	return nil
}

func (a *Assertion) checkBodyArrayExist(body interface{}) error {
	return nil
}

func (a *Assertion) checkStatusExact(status int, expected int) error {
	if status != expected {
		return fmt.Errorf("on::status expected::%d , but receive::%d", expected, status)
	}
	return nil
}

func (a *Assertion) checkStatusNot(status int, expected int) error {
	if status == expected {
		return fmt.Errorf("on::status expected::not %d , but receive::%d", expected, status)
	}
	return nil
}

func (a *Assertion) checkStatus(result Result) error {
	switch a.Check {
	case "exact":
		var num int
		switch a.Expected.(type) {
		case int:
			num = a.Expected.(int)
		case float64:
			num = int(a.Expected.(float64))
		case string:
			num = utils.StringToInt(a.Expected.(string))
		}
		return a.checkStatusExact(result.Status, int(num))
	case "not":
		return a.checkStatusNot(result.Status, a.Expected.(int))
	default:
		return fmt.Errorf("on::status command::%s is no supported", a.Check)
	}
}

func (a *Assertion) checkBody(result Result) error {
	switch a.Check {
	case "exact":
		return a.checkBodyExact(result.Body)
	case "type":
		return a.checkBodyType(result.Body)
	case "exist":
		return a.checkBodyExist(result.Body)
	default:
		return fmt.Errorf("on::body command::%s is no supported", a.Check)
	}
}

func (a *Assertion) checkBodyExact(body interface{}) error {
	expectedBody := a.Expected.(map[string]interface{})
	expectedFlattenBody := utils.Flatten(expectedBody)
	actualFlattenBody := utils.Flatten(body.(map[string]interface{}))
	err := utils.CompareMap(expectedFlattenBody, actualFlattenBody)
	if err != nil {
		return fmt.Errorf("on::body command::exact %s", err.Error())
	}
	return nil
}

// Get value from At property
// Replace the value part of the expected body with the value from At property
// For example:
// expected assertion: array::type::@children
// At: {"children": "string"}
// change to: array::type::string
func (a *Assertion) getArrayTypeCheckValue(raw interface{}) interface{} {
	rawString := utils.InterfaceToString(raw)
	arr := utils.SplitBy(rawString, "::")
	hasRef := len(arr) == 3 && utils.ContainerSubString(rawString, "@")
	if hasRef {
		arr2WithoutAt := utils.SplitBy(arr[2], "@")
		return fmt.Sprintf("%s::%s::%s", arr[0], arr[1], a.At[arr2WithoutAt[1]])
	}

	return raw
}

func (a *Assertion) checkBodyType(body interface{}) error {
	expectedBody := a.Expected.(map[string]interface{})
	expectedFlattenBody := utils.Flatten(expectedBody)
	actualFlattenBody := utils.Flatten(body.(map[string]interface{}))
	var typeValidator validators.IValidator = validators.NewTypeValidator()
	var errs []error

	for k, v := range expectedFlattenBody {
		isArrayTypeCheck := utils.ContainerSubString(utils.InterfaceToString(v), "array")
		if isArrayTypeCheck {
			v = a.getArrayTypeCheckValue(v)
		}
		err := typeValidator.Validate(v, actualFlattenBody[k])
		if err != nil {
			errs = append(errs, fmt.Errorf("key::%s %s ;", k, err.Error()))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("on::body command::type %s", errs)
	}

	return nil
}

func (a *Assertion) checkBodyExist(body interface{}) error {
	expectedBody := a.Expected.(map[string]interface{})
	expectedFlattenBody := utils.Flatten(expectedBody)
	actualFlattenBody := utils.Flatten(body.(map[string]interface{}))
	var typeValidator validators.IValidator = validators.NewExistValidator()

	err := typeValidator.Validate(expectedFlattenBody, actualFlattenBody)
	if err != nil {
		return fmt.Errorf("on::body command::exist %s", err.Error())
	}

	return nil
}

func (a *Assertion) checkHeader(result Result) error {
	switch a.Check {
	case "exist":
		return a.checkHeaderExist(result.Headers)
	}

	return nil
}

func (a *Assertion) checkHeaderExist(headers map[string]string) error {
	expectedHeaders := a.Expected.(map[string]interface{})
	expectedFlattenHeaders := utils.Flatten(expectedHeaders)
	actualFlattenHeaders := utils.Flatten(utils.MapStringStringToMapStringInterface(headers))
	var typeValidator validators.IValidator = validators.NewExistValidator()

	err := typeValidator.Validate(expectedFlattenHeaders, actualFlattenHeaders)
	if err != nil {
		return fmt.Errorf("on::header command::exist %s", err.Error())
	}

	return nil
}
