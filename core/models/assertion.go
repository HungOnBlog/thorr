package models

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/validators"
	"github.com/HungOnBlog/thorr/utils"
)

type Assertion struct {
	On       string      `json:"on" yaml:"on"`
	Check    string      `json:"check" yaml:"check"`
	Expected interface{} `json:"expected" yaml:"expected"`
}

func (a *Assertion) CheckAssertion(result Result) error {
	switch a.On {
	case "status":
		return a.checkStatus(result)
	case "body":
		return a.checkBody(result)
	case "header":
		return a.checkHeader(result)
	default:
		return fmt.Errorf("on::%s is no supported", a.On)
	}
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
		return a.checkStatusExact(result.Status, a.Expected.(int))
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

func (a *Assertion) checkBodyType(body interface{}) error {
	expectedBody := a.Expected.(map[string]interface{})
	expectedFlattenBody := utils.Flatten(expectedBody)
	actualFlattenBody := utils.Flatten(body.(map[string]interface{}))
	var typeValidator validators.IValidator = validators.NewTypeValidator()
	var errs []error

	for k, v := range expectedFlattenBody {
		err := typeValidator.Validate(v, actualFlattenBody[k])
		if err != nil {
			errs = append(errs, err)
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
