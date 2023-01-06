package utils

import "testing"

func TestFlatten(t *testing.T) {
	rawJson := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key3": "value3",
			"key4": map[string]interface{}{
				"key5": "value5",
				"key6": map[string]interface{}{
					"key7": "value7",
					"key8": map[string]interface{}{
						"key9": 1,
						"key10": map[string]interface{}{
							"key11": 1,
						},
					},
				},
			},
		},
	}

	flatJson := FlattenJson(rawJson)

	if flatJson["key1"] != "value1" {
		t.Error("Expected value1 for key1")
	}

	if flatJson["key2.key3"] != "value3" {
		t.Error("Expected value3 for key2.key3")
	}

	if flatJson["key2.key4.key5"] != "value5" {
		t.Error("Expected value5 for key2.key4.key5")
	}

	if flatJson["key2.key4.key6.key7"] != "value7" {
		t.Error("Expected value7 for key2.key4.key6.key7")
	}

	if flatJson["key2.key4.key6.key8.key9"] != 1 {
		t.Error("Expected value9 for key2.key4.key6.key8.key9")
	}

	if flatJson["key2.key4.key6.key8.key10.key11"] != 1 {
		t.Error("Expected value11 for key2.key4.key6.key8.key10.key11")
	}
}
