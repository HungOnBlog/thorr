package main

import (
	"fmt"

	"github.com/HungOnBlog/thorr/infrastructure/parser"
)

func main() {
	filePath := "./tests/integrations/template.yaml"
	parser := parser.NewYamlParser()
	suit, err := parser.Parse(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(suit)
}
