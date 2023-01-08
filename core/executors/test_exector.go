package executors

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/infrastructure/requester"
)

type TestExecutor struct{}

func NewTestExecutor() *TestExecutor {
	return &TestExecutor{}
}

func (e *TestExecutor) Execute(test *models.Test) error {
	fmt.Println("Running test: ", test.Name)
	var requester requester.IRequester = requester.NewHttpRequester()
	result, err := requester.DoRequest(*test)
	if err != nil {
		return err
	}

	var errs []error
	for _, assertion := range test.Assertions {
		err := assertion.CheckAssertion(result)
		if err != nil {
			errs = append(errs, err)
		} else {
			fmt.Printf("✅ test::%v assertion::%v passed \n", test.Name, assertion)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("❌ Error: %v", errs)
	}

	return nil
}
