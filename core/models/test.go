package models

type Test struct {
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Request     TestRequest            `json:"request" yaml:"request"`
	Assertions  []Assertion            `json:"assertions" yaml:"assertions"`
	Variables   map[string]interface{} `json:"variables" yaml:"variables"`
}

type TestRequest struct {
	BaseURL    string                 `json:"base_url" yaml:"base_url"`
	Path       string                 `json:"path" yaml:"path"`
	PathParams map[string]interface{} `json:"path_params" yaml:"path_params"`
	Method     string                 `json:"method" default:"GET" yaml:"method"`
	Header     map[string]interface{} `json:"header" yaml:"header"`
	Body       map[string]interface{} `json:"body" yaml:"body"`
	Query      map[string]interface{} `json:"query" yaml:"query"`
}
