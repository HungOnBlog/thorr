package models

type TestSuit struct {
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	Status      string          `json:"status" default:"created" yaml:"status"`
	Default     TestSuitDefault `json:"default" yaml:"default"`
	Tests       []Test          `json:"tests" yaml:"tests"`
}

type TestSuitDefault struct {
	BaseURL string            `json:"base_url" yaml:"base_url"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}
