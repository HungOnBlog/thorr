package utils

import "testing"

func TestIsString(t *testing.T) {
	if !IsString("test") {
		t.Errorf("TestIsString: expected true, got false")
	}

	if IsString(1) {
		t.Errorf("TestIsString: expected false, got true")
	}
}

func TestIsInt(t *testing.T) {
	if !IsInt(1) {
		t.Errorf("TestIsInt: expected true, got false")
	}

	if IsInt("test") {
		t.Errorf("TestIsInt: expected false, got true")
	}
}

func TestIsFloat(t *testing.T) {
	if !IsFloat(1.0) {
		t.Errorf("TestIsFloat: expected true, got false")
	}

	if IsFloat("test") {
		t.Errorf("TestIsFloat: expected false, got true")
	}
}

func TestIsBool(t *testing.T) {
	if !IsBool(true) {
		t.Errorf("TestIsBool: expected true, got false")
	}

	if IsBool("test") {
		t.Errorf("TestIsBool: expected false, got true")
	}
}

func TestIsArray(t *testing.T) {
	if !IsArray([]interface{}{}) {
		t.Errorf("TestIsArray: expected true, got false")
	}

	if IsArray("test") {
		t.Errorf("TestIsArray: expected false, got true")
	}

	if IsArray(map[string]interface{}{}) {
		t.Errorf("TestIsArray: expected false, got true")
	}
}

func TestIsMap(t *testing.T) {
	if !IsMap(map[string]interface{}{}) {
		t.Errorf("TestIsMap: expected true, got false")
	}

	if IsMap("test") {
		t.Errorf("TestIsMap: expected false, got true")
	}

	if IsMap([]interface{}{}) {
		t.Errorf("TestIsMap: expected false, got true")
	}
}

func TestIsDateString(t *testing.T) {
	if !IsDateString("2019-01-01") {
		t.Errorf("TestIsDateString: expected true, got false")
	}

	if IsDateString("2019-01-01 00:00:00") {
		t.Errorf("TestIsDateString: expected false, got true")
	}

	if IsDateString("00:00:00") {
		t.Errorf("TestIsDateString: expected false, got true")
	}

	if IsDateString(1) {
		t.Errorf("TestIsDateString: expected false, got true")
	}
}

func TestIsDateTimeString(t *testing.T) {
	if !IsDateTimeString("2019-01-01 00:00:00") {
		t.Errorf("TestIsDateTimeString: expected true, got false")
	}

	if IsDateTimeString("2019-01-01") {
		t.Errorf("TestIsDateTimeString: expected false, got true")
	}

	if IsDateTimeString("00:00:00") {
		t.Errorf("TestIsDateTimeString: expected false, got true")
	}

	if IsDateTimeString(1) {
		t.Errorf("TestIsDateTimeString: expected false, got true")
	}
}

func TestIsTimeString(t *testing.T) {
	if !IsTimeString("00:00:00") {
		t.Errorf("TestIsTimeString: expected true, got false")
	}

	if IsTimeString("2019-01-01") {
		t.Errorf("TestIsTimeString: expected false, got true")
	}

	if IsTimeString("2019-01-01 00:00:00") {
		t.Errorf("TestIsTimeString: expected false, got true")
	}

	if IsTimeString(1) {
		t.Errorf("TestIsTimeString: expected false, got true")
	}
}
