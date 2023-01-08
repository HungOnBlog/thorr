package parser

import (
	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/utils"
)

type YamlParser struct{}

func NewYamlParser() YamlParser {
	return YamlParser{}
}

func (j *YamlParser) Parse(path string) (models.TestSuit, error) {
	fileContent, err := utils.ReadFile(path)
	if err != nil {
		return models.TestSuit{}, err
	}

	var suit models.TestSuit
	err = utils.UnmarshalYaml(fileContent, &suit)
	if err != nil {
		return models.TestSuit{}, err
	}

	return suit, nil
}
