package main

import (
	"github.com/HungOnBlog/thorr/infrastructure/cli"
)

func main() {
	cliApp := cli.NewThorCliApp()
	err := cliApp.Run()
	if err != nil {
		panic(err)
	}
}
