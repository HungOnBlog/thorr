package models

type TestSuit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status" default:"created"`
	Tests       []Test `json:"tests"`
}
