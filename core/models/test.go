package models

type Test struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Request     TestRequest `json:"request"`
	Assertions  []Assertion `json:"assertions"`
}

type TestRequest struct {
	BaseURL    string                 `json:"base_url"`
	Path       string                 `json:"path"`
	PathParams map[string]interface{} `json:"path_params"`
	Method     string                 `json:"method" default:"GET"`
	Header     map[string]interface{} `json:"header"`
	Body       map[string]interface{} `json:"body"`
	Query      map[string]interface{} `json:"query"`
}

type TestExpected struct {
	Status int                    `json:"status"` // 200, 404, etc
	Body   map[string]interface{} `json:"body"`
}
