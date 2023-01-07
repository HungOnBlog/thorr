package utils

import (
	"strings"
	"testing"
)

func TestSplitBy(t *testing.T) {
	str := "1,2,3,4"
	arr := SplitBy(str, ",")
	if len(arr) != 4 {
		t.Errorf("TestSplitBy: expected 4, got %v", len(arr))
	}

	for _, v := range arr {
		// check if value contains space (not trimmed)
		if strings.Contains(v, " ") {
			t.Errorf("TestSplitBy: expected no space, got %v", v)
		}
	}
}
