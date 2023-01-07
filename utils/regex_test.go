package utils

import "testing"

func TestMathRegex(t *testing.T) {
	emailRegex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	if !MatchRegex(emailRegex, "example@gmail.com") {
		t.Errorf("TestMathRegex: expected true, got false")
	}

	if MatchRegex(emailRegex, "example@gmail") {
		t.Errorf("TestMathRegex: expected false, got true")
	}

	if MatchRegex(emailRegex, "example@gmail.") {
		t.Errorf("TestMathRegex: expected false, got true")
	}
}
