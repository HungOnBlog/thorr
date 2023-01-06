package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
		err := test.Execute()
		if err != nil {
			errors <- err
		}
	}

	return nil
}

func (t *Test) Execute() error {
	fmt.Println("Executing test: " + t.Name)
	return nil
}
