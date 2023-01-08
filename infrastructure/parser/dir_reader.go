package parser

import (
	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/utils"
)

type DirParser struct{}

func NewDirParser() *DirParser {
	return &DirParser{}
}

func (d *DirParser) Parse(dirPath string) ([]models.TestSuit, error) {
	filePaths, err := utils.ListAllFilesPathInDir(dirPath)
	if err != nil {
		return nil, err
	}

	var testSuits []models.TestSuit
	var parser IParse
	for _, filePath := range filePaths {
		if utils.IsJsonFile(filePath) {
			parser = &JsonParser{}
		} else if utils.IsYamlFile(filePath) {
			parser = &YamlParser{}
		} else {
			continue
		}

		testSuit, err := parser.Parse(filePath)
		if err != nil {
			return nil, err
		}

		testSuits = append(testSuits, testSuit)
	}

	return testSuits, nil
}
