package utils

import "testing"

func TestFlatten(t *testing.T) {
	deepObj := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": "baz",
			"inside": map[string]interface{}{
				"deep": "value",
			},
		},
	}

	flatObj := Flatten(deepObj)

	if flatObj["foo.bar"] != "baz" {
		t.Errorf("Expected foo.bar to be baz, got %v", flatObj["foo.bar"])
	}

	if flatObj["foo.inside.deep"] != "value" {
		t.Errorf("Expected foo.inside.deep to be value, got %v", flatObj["foo.inside.deep"])
	}
}
