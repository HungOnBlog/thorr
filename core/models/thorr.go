package models

import (
	"fmt"

	"github.com/HungOnBlog/thorr/common/utils"
)

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
		return t.runIntegrationTests()
	case "load":
		return t.runLoadTests()
	default:
		panic("Unknown command")
	}
}

func (t *Thorr) Stop() error {
	return nil
}

func (t *Thorr) runIntegrationTests() error {
	testSuites, err := t.testFinder()
	if err != nil {
		return err
	}

	if len(*testSuites) == 0 {
		return fmt.Errorf("no test suites found")
	}

	errors := make(chan error, len(*testSuites))

	for _, testSuite := range *testSuites {
		err := testSuite.Execute()
		if err != nil {
			errors <- err
		}
	}

	return nil
}

func (t *Thorr) runLoadTests() error {
	return nil
}

func (t *Thorr) testFinder() (*[]TestSuit, error) {
	path := t.Options.File
	isFile := utils.IsFile(path)
	var testSuits []TestSuit

	if isFile {
		content := utils.ReadJsonFile(path)

		testSuit, err := UnmarshalTestSuit(content)
		if err != nil {
			return nil, err
		}

		testSuits = append(testSuits, testSuit)
	} else {
		files := utils.LoadAllFilePathsIn(path)
		for _, file := range files {
			content := utils.ReadJsonFile(file)

			testSuit, err := UnmarshalTestSuit(content)
			if err != nil {
				return nil, err
			}

			testSuits = append(testSuits, testSuit)
		}
	}

	return &testSuits, nil
}
