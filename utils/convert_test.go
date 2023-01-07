package utils

import (
	"testing"
	"time"
)

func TestMapStringInterfaceToMapStringString(t *testing.T) {
	raw := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	result := MapStringInterfaceToMapStringString(raw)

	if len(result) != len(expected) {
		t.Errorf("TestMapStringInterfaceToMapStringString: expected %v, got %v", len(expected), len(result))
	}
}

func TestInterfaceToString(t *testing.T) {
	var i interface{} = "test"
	s := InterfaceToString(i)
	if s != "test" {
		t.Errorf("TestInterfaceToString: expected %v, got %v", "test", s)
	}
}

func TestStringToInt(t *testing.T) {
	s := "123"
	i := StringToInt(s)
	if i != 123 {
		t.Errorf("TestStringToInt: expected %v, got %v", 123, i)
	}
}

func TestStringToFloat(t *testing.T) {
	s := "123.456"
	f := StringToFloat(s)
	if f != 123.456 {
		t.Errorf("TestStringToFloat64: expected %v, got %v", 123.456, f)
	}
}

func TestStringToDate(t *testing.T) {
	s := "2018-01-01"
	d, err := StringToDate(s)
	if err != nil {
		t.Errorf("TestStringToDate: expected %v, got %v", nil, err)
	}
	if d.Year() != 2018 || d.Month() != 1 || d.Day() != 1 {
		t.Errorf("TestStringToDate: expected %v, got %v", "2018-01-01", d)
	}
}

func TestStringToDateWithFormat(t *testing.T) {
	format := "2022/01/01"
	d, err := StringToDateWithFormat(format, "2006/01/02")
	if err != nil {
		t.Errorf("TestStringToDateWithFormat: expected %v, got %v", nil, err)
	}
	if d.Year() != 2022 || d.Month() != 1 || d.Day() != 1 {
		t.Errorf("TestStringToDateWithFormat: expected %v, got %v", "2022/01/01", d)
	}
}

func TestStringToUTC(t *testing.T) {
	s := "2021-01-01T00:00:00Z"
	d, err := StringToUTC(s)
	if err != nil {
		t.Errorf("TestStringUTC: expected %v, got %v", nil, err)
	}
	if d.Year() != 2021 || d.Month() != 1 || d.Day() != 1 {
		t.Errorf("TestStringUTC: expected %v, got %v", "2021-01-01T00:00:00Z", d)
	}
}

func TestStringToTime(t *testing.T) {
	// 2021-01-01T00:00:00Z
	s := "2021-01-01T00:00:00Z"
	d, err := StringToTime(s)
	if err != nil {
		t.Errorf("TestStringToTime: expected %v, got %v", nil, err)
	}

	if d.Year() != 2021 || d.Month() != 1 || d.Day() != 1 {
		t.Errorf("TestStringToTime: expected %v, got %v", "2021-01-01T00:00:00Z", d)
	}
}

func TestStringToTimeWithFormat(t *testing.T) {
	s := "2021-01-01T00:00:00Z"
	d, err := StringToTimeWithFormat(s, time.RFC3339)
	if err != nil {
		t.Errorf("TestStringToTimeWithFormat: expected %v, got %v", nil, err)
	}

	if d.Year() != 2021 || d.Month() != 1 || d.Day() != 1 {
		t.Errorf("TestStringToTimeWithFormat: expected %v, got %v", "2021-01-01T00:00:00Z", d)
	}
}

func TestStringUnixToTime(t *testing.T) {
	s := "1612086400"
	d, err := StringUnixToTime(s)
	if err != nil {
		t.Errorf("TestStringUnixToTime: expected %v, got %v", nil, err)
	}

	if d.Year() != 2021 || d.Month() != 1 || d.Day() != 31 {
		t.Errorf("TestStringUnixToTime: expected %v, got %v", "2021-01-31T00:00:00Z", d)
	}
}
