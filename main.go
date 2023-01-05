package main

import (
	simplecli "github.com/HungOnBlog/thorr/app/simple_cli"
)

func main() {
	thorSimpleCli := &simplecli.ThorrSimpleCli{}
	thorSimpleCli.New()
	thorSimpleCli.Run()
}
