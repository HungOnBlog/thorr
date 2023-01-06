package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"os"

	"github.com/HungOnBlog/thorr/common/requester"
	"github.com/HungOnBlog/thorr/common/utils"
)

type TestSuit struct {
	Id          string `json:"id" gorm:"PrimaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BaseURL     string `json:"baseURL"`
	Tests       []Test `json:"tests"`
}

type Test struct {
	Id          string                 `json:"id" gorm:"PrimaryKey"`
	SuitId      string                 `json:"suitId"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`
	BaseURL     string                 `json:"baseURL"`
	Path        string                 `json:"path"`
	Body        map[string]interface{} `json:"body" gorm:"type:jsonb;not null;default:'{}'::jsonb"`
	Headers     map[string]interface{} `json:"headers" gorm:"type:jsonb;not null;default:'{}'::jsonb"`
	Params      map[string]interface{} `json:"params" gorm:"type:jsonb;not null;default:'{}'::jsonb"`
	Expected    TestExpected           `json:"expected" gorm:"type:jsonb;not null;default:'{}'::jsonb"`
}

type TestExpected struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

func (t *TestExpected) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *TestExpected) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), t)
}

func MarshalTestSuit(testSuit TestSuit) ([]byte, error) {
	return json.Marshal(testSuit)
}

func UnmarshalTestSuit(data []byte) (TestSuit, error) {
	var testSuit TestSuit
	err := json.Unmarshal(data, &testSuit)
	return testSuit, err
}

func (t *TestSuit) Execute() error {
	fmt.Println("Executing test suit: " + t.Name)
	numberOfTests := len(t.Tests)
	fmt.Println("Number of tests: ", numberOfTests)
	errors := make(chan error, numberOfTests)

	for _, test := range t.Tests {
		test.BaseURL = t.BaseURL
		err := test.Execute()
		if err != nil {
			errors <- err
		}
	}

	return nil
}

func (t *Test) Execute() error {
	fmt.Println("Executing test: " + t.Name)
	httpRequester := requester.NewHttpRequester()
	status, body, err := httpRequester.Request(t.Method, t.BaseURL+t.Path, t.Headers, t.Params, t.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Fprintln(os.Stdout, "Status: ", status)
	fmt.Fprintln(os.Stdout, "Body: ", utils.FlattenJson(utils.JsonToMap(body)))

	isValid, err := t.validate(status, body)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	if !isValid {
		fmt.Println("Test failed ☠️")
	} else {
		fmt.Println("Test passed ✅")
	}
	return nil
}

func (t *Test) validate(status int, body []byte) (bool, error) {
	if t.Expected.Status != status {
		return false, fmt.Errorf("expected status %d, got %d", t.Expected.Status, status)
	}

	flatResponseBody := utils.FlattenJson(utils.JsonToMap(body))
	flatExpectedBody := utils.FlattenJson(t.Expected.Body)

	for key, value := range flatExpectedBody {
		if flatResponseBody[key] != value {
			return false, fmt.Errorf("expected %s to be %s, got %s ❌", key, value, flatResponseBody[key])
		}
	}

	return true, nil
}
