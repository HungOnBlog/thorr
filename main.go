package main

import (
	"fmt"

	"github.com/HungOnBlog/thorr/infrastructure/parser"
)

func main() {
	folderPath := "./tests/integrations"
	dirParser := parser.NewDirParser()
	suits, err := dirParser.Parse(folderPath)
	if err != nil {
		panic(err)
	}

	for _, suit := range suits {
		fmt.Println(suit)
	}
}
