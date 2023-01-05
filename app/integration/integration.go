package integration

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
)

func LoadTestSuits(path string) []models.TestSuit {
	if path == "" {
		fmt.Println("No test suit file or directory is provided. 💁")
		return []models.TestSuit{}
	}
	return []models.TestSuit{}
}
