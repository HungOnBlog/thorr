package models

type TestSuit struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Status      string          `json:"status" default:"created"`
	BaseURL     string          `json:"base_url"`
	Default     TestSuitDefault `json:"default"`
	Tests       []Test          `json:"tests"`
}

type TestSuitDefault struct {
	BaseURL string            `json:"base_url"`
	Headers map[string]string `json:"headers"`
}
