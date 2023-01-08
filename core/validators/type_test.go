package validators

import (
	"testing"
)

func TestTypeValidator(t *testing.T) {
	validator := NewTypeValidator()
	cases := []struct {
		expected interface{}
		actual   interface{}
		isNotErr bool
	}{
		{"string", "test", true},
		{"string", 1, false},
		{"int", 1, true},
		{"int", "test", false},
		{"float", 1.0, true},
		{"float", "test", false},
		{"bool", true, true},
		{"bool", "test", false},
		{"date", "2018-01-01", true},
		{"date", "00:00:00", false},
		{"date", "test", false},
		{"time", "00:00:00", true},
		{"time", "2018-01-01", false},
		{"time", "test", false},
		{"string::email", "", false},
		{"string::email", "test", false},
		{"string::email", "example.com", false},
		{"string::email", "example@email.com", true},
		{"string::length::5-10", "12345", true},
		{"string::length::5-10", "1234567890", true},
		{"string::length::5-10", "12345678901", false},
		{"string::length::5-10", "1234", false},
		{"string::uuid", "a0b1ac03-fa14-4691-a6ca-44d68b604611", true},
		{"string::uuid", "a0b1ac03-fa14-4691-a6ca-44d68b60461", false},
		{"string::base64", "aGVsbG8gd29ybGQ=", true},
		{"string::base64", "aGVsbG8gd29ybGQ", false},
		{"string::regex::^\\d+$", "asdf", false},
		{"string::regex::^\\d+$", "1234", true},
		{"int::range::1-10", 1, true},
		{"int::range::1-10", 10, true},
		{"int::range::1-10", 0, false},
		{"int::range::1-10", 11, false},
		{"int::lt::10", 9, true},
		{"int::lt::10", 10, false},
		{"int::lt::10", 11, false},
		{"int::lte::10", 9, true},
		{"int::lte::10", 10, true},
		{"int::lte::10", 11, false},
		{"int::gt::10", 9, false},
		{"int::gt::10", 10, false},
		{"int::gt::10", 11, true},
		{"int::gte::10", 9, false},
		{"int::gte::10", 10, true},
		{"int::gte::10", 11, true},
		{"float::range::1.0-10.0", 1.0, true},
		{"float::range::1.0-10.0", 10.0, true},
		{"float::range::1.0-10.0", 0.0, false},
		{"float::range::1.0-10.0", 11.0, false},
		{"float::lt::10.0", 9.0, true},
		{"float::lt::10.0", 10.0, false},
		{"float::lt::10.0", 11.0, false},
		{"float::lte::10.0", 9.0, true},
		{"float::lte::10.0", 10.0, true},
		{"float::lte::10.0", 11.0, false},
		{"float::gt::10.0", 9.0, false},
		{"float::gt::10.0", 10.0, false},
		{"float::gt::10.0", 11.0, true},
		{"float::gte::10.0", 9.0, false},
		{"float::gte::10.0", 10.0, true},
		{"float::gte::10.0", 11.0, true},
		{"date::before::2018-01-01", "2017-12-31", true},
		{"date::before::2018-01-01", "2018-01-01", false},
		{"date::before::2018-01-01", "2018-01-02", false},
		{"date::after::2018-01-01", "2017-12-31", false},
		{"date::after::2018-01-01", "2018-01-01", false},
		{"date::after::2018-01-01", "2018-01-02", true},
		{"date::range::2018-01-01 to 2018-01-31", "2018-01-01", true},
		{"date::range::2018-01-01 to 2018-01-31", "2018-01-31", true},
		{"date::range::2018-01-01 to 2018-01-31", "2018-01-02", true},
		{"date::range::2018-01-02 to 2018-01-31", "2018-01-01", false},
		{"date::range::2018-01-01 to 2018-01-31", "2018-02-01", false},
		{"date::range::2018-01-01 to 2018-01-30", "2018-01-31", false},
		{"date::utc", "2010-11-12T13:14:15Z00:00", false},
		{"date::utc", "2006-01-02T15:04:05Z07:00", false},
		{"date::utc", "2010-11-12T13:14:15+01:00", true},
		{"date::utc", "2010-11-12T13:14:15-01:00", true},
		{"date::format::2006-01-02", "2018-01-01", true},
		{"date::format::2006-01-02", "2018-01-01 00:00:00", false},
		{"date::format::2006/01/02", "2018-01-01", false},
		{"date::format::2006/01/02", "2018-01-01 00:00:00", false},
		{"date::format::2006-01-02 15:04:05", "2018-01-01 00:00:00", true},
		{"date::iso8601", "2018-01-01", false},
		{"date::iso8601", "2018-01-01 00:00:00", false},
		{"date::iso8601", "2018-01-01T00:00:00Z", true},
		{"date::iso8601", "2018-01-01T00:00:00+01:00", true},
	}

	for _, c := range cases {
		err := validator.Validate(c.expected, c.actual)
		if c.isNotErr && err != nil {
			t.Errorf("Validate(%v, %v) is error: %v", c.expected, c.actual, err)
		}
	}
}
