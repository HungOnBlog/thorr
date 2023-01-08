package main

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/infrastructure/requester"
	"github.com/HungOnBlog/thorr/utils"
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
					Path:    "/anything/hello",
					Method:  "GET",
					Header:  map[string]interface{}{"Content-Type": "application/json"},
				},
				Assertions: []models.Assertion{
					{
						On:       "status",
						Check:    "exact",
						Expected: 200,
					},
				},
			},
		},
	}

	httpRequester := requester.NewHttpRequester()
	for _, test := range suit.Tests {
		result, err := httpRequester.DoRequest(test)
		if err != nil {
			panic(err)
		}

		fmt.Println(result.Status)
		flattenMap := utils.Flatten(result.Body.(map[string]interface{}))
		for k, v := range flattenMap {
			fmt.Println(k, v)
		}
	}
}
