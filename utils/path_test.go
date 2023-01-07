package utils

import "testing"

func TestGetPlaceHolders(t *testing.T) {
	paths := []string{
		"/users/:id",
		"/users/:id/:name",
		"/users/:id/phone/:phone_id",
	}

	expected := [][]string{
		{":id"},
		{":id", ":name"},
		{":id", ":phone_id"},
	}

	for i, v := range paths {
		placeholders := GetPathPlaceHolders(v)
		if len(placeholders) != len(expected[i]) {
			t.Errorf("TestGetPlaceHolders: expected %v, got %v", expected[i], placeholders)
		}
	}
}

func TestGenPaths(t *testing.T) {
	paths := []string{
		"/users/:id",
		"/users/:id/:name",
		"/users/:id/phone/:phone_id",
		"/users/:id/phone/:phone_id/:name",
	}

	expected := [][]string{
		{"/users/1", "/users/2"},
		{"/users/1/John", "/users/2/John"},
		{"/users/1/phone/1", "/users/1/phone/2", "/users/2/phone/1", "/users/2/phone/2"},
		{"/users/1/phone/1/John", "/users/1/phone/2/John", "/users/2/phone/1/John", "/users/2/phone/2/John"},
	}

	for i, v := range paths {
		placeholders := GetPathPlaceHolders(v)
		paths := GeneratePaths(v, placeholders, map[string][]string{
			"id":       {"1", "2"},
			"name":     {"John"},
			"phone_id": {"1", "2"},
		})

		if len(paths) != len(expected[i]) {
			t.Errorf("TestGenPaths: expected %v, got %v", expected[i], paths)
		}
	}
}
