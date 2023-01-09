package executors

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/utils"
)

type TestSuitExecutor struct{}

func NewTestSuitExecutor() *TestSuitExecutor {
	return &TestSuitExecutor{}
}

func (e *TestSuitExecutor) Execute(testSuit *models.TestSuit) error {
	var errs []error
	var globalVariables map[string]interface{} = make(map[string]interface{})
	textExecutor := NewTestExecutor()
	for _, test := range testSuit.Tests {
		test.Request.BaseURL = testSuit.Default.BaseURL
		test.Request.Header = utils.MapAssignAll(
			test.Request.Header,
			utils.MapStringStringToMapStringInterface(testSuit.Default.Headers),
		)

		err := textExecutor.Execute(&test, &globalVariables)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		fmt.Println("❌ Error: ", errs)
	}

	return nil
}
