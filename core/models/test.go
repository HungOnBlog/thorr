package models

type Test struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Request     TestRequest  `json:"request"`
	Expected    TestExpected `json:"expected"`
}

type TestRequest struct {
	BaseURL string                 `json:"baseURL"`
	Method  string                 `json:"method" default:"GET"`
	Header  map[string]interface{} `json:"header"`
	Body    map[string]interface{} `json:"body"`
	Query   map[string]interface{} `json:"query"`
}

type TestExpected struct {
	Status int                    `json:"status"` // 200, 404, etc
	Body   map[string]interface{} `json:"body"`
}
