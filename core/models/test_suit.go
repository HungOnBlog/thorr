package models

import "encoding/json"

type TestSuit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	BaseURL     string `json:"baseURL"`
	Tests       []Test `json:"tests"`
}

type Test struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Body        interface{} `json:"body"`
	Headers     interface{} `json:"headers"`
	Params      interface{} `json:"params"`
	Expected    interface{} `json:"expected"`
}

func MarshalTestSuit(testSuit TestSuit) ([]byte, error) {
	return json.Marshal(testSuit)
}

func UnmarshalTestSuit(data []byte) (TestSuit, error) {
	var testSuit TestSuit
	err := json.Unmarshal(data, &testSuit)
	return testSuit, err
}
