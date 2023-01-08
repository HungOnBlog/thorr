package validators

import (
	"fmt"
)

type ExistValidator struct{}

func NewExistValidator() *ExistValidator {
	return &ExistValidator{}
}

func (v *ExistValidator) Validate(expected interface{}, actual interface{}) error {
	flattenExpected := expected.(map[string]interface{})
	flattenActual := actual.(map[string]interface{})
	var errs []error
	for k, v := range flattenExpected {
		expectedExist := v.(bool)
		if expectedExist {
			if _, ok := flattenActual[k]; !ok {
				errs = append(errs, fmt.Errorf("expected::%v exists but actual not exist", k))
			}
		} else {
			if _, ok := flattenActual[k]; ok {
				errs = append(errs, fmt.Errorf("expected::%v not exists but actual exists", k))
			}
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("on::body command::exist %s", errs)
	}

	return nil
}
