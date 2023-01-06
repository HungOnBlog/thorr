package main

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/infrastructure/requester"
)

func main() {
	suit := models.TestSuit{
		Name:        "Test Suit",
		Description: "Test Suit Description",
		Status:      "created",
		Tests: []models.Test{
			{
				Name:        "Test 1",
				Description: "Test 1 Description",
				Request: models.TestRequest{
					BaseURL: "https://httpbin.org",
					Path:    "/uuid",
					Method:  "GET",
					Header:  map[string]interface{}{"Content-Type": "application/json"},
				},
				Expected: models.TestExpected{
					Status: 200,
					Body:   map[string]interface{}{"name": "Hung"},
				},
			},
		},
	}

	httpRequester := requester.NewHttpRequester()
	for _, test := range suit.Tests {
		status, res, err := httpRequester.DoRequest(test)
		if err != nil {
			panic(err)
		}

		fmt.Println(status)
		fmt.Println(res)
	}
}
