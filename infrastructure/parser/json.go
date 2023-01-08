package parser

import (
	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/utils"
)

type JsonParser struct{}

func NewJsonParser() JsonParser {
	return JsonParser{}
}

func (j *JsonParser) Parse(path string) (models.TestSuit, error) {
	fileContent, err := utils.ReadFile(path)
	if err != nil {
		return models.TestSuit{}, err
	}

	var suit models.TestSuit
	err = utils.UnmarshalJson(fileContent, &suit)
	if err != nil {
		return models.TestSuit{}, err
	}

	return suit, nil
}
