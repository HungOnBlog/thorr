package models

import "encoding/json"

type TestSuit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BaseURL     string `json:"baseURL"`
	Tests       []Test `json:"tests"`
}

type Test struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`
	Path        string                 `json:"path"`
	Body        map[string]interface{} `json:"body"`
	Headers     map[string]interface{} `json:"headers"`
	Params      map[string]interface{} `json:"params"`
	Expected    TestExpected           `json:"expected"`
}

type TestExpected struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

func MarshalTestSuit(testSuit TestSuit) ([]byte, error) {
	return json.Marshal(testSuit)
}

func UnmarshalTestSuit(data []byte) (TestSuit, error) {
	var testSuit TestSuit
	err := json.Unmarshal(data, &testSuit)
	return testSuit, err
}
