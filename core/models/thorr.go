package models

import "fmt"

type ThorrOptions struct {
	File    string `json:"file"`
	Spawns  int    `json:"spawns" default:"1"`
	Command string `json:"command"`
}

type Thorr struct {
	Options ThorrOptions `json:"options"`
}

func NewThorr(Options ThorrOptions) *Thorr {
	return &Thorr{
		Options: Options,
	}
}

func (t *Thorr) Start() error {
	fmt.Println("Starting thorr")
	switch t.Options.Command {
	case "integration":
		fmt.Println("Starting integration test")
	case "load":
		fmt.Println("Starting load test")
	default:
		panic("Unknown command")
	}

	return nil
}
