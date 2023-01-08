package thorr

import (
	"fmt"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/HungOnBlog/thorr/infrastructure/parser"
	"github.com/HungOnBlog/thorr/utils"
)

type ThorrApp struct {
	FilePath string
	Spawn    int
	RunType  string
}

func NewThorrApp(filePath string, spawn int, runType string) *ThorrApp {
	return &ThorrApp{
		FilePath: filePath,
		Spawn:    spawn,
		RunType:  runType,
	}
}

func (t *ThorrApp) Run() error {
	switch t.RunType {
	case "integration":
		return t.runIntegration()
	case "load":
		return t.runLoad()
	default:
		return fmt.Errorf("command not found")
	}
}

func (t *ThorrApp) getTestSuits() ([]models.TestSuit, error) {
	isFile := utils.IsFile(t.FilePath)
	if isFile {
		return t.getTestSuitsFromFile()
	}

	return t.getTestSuitsFromDir()
}

func (t *ThorrApp) getTestSuitsFromFile() ([]models.TestSuit, error) {
	var p parser.IParse
	if utils.IsJsonFile(t.FilePath) {
		p = &parser.JsonParser{}
	} else if utils.IsYamlFile(t.FilePath) {
		p = &parser.YamlParser{}
	} else {
		return nil, fmt.Errorf("file type not supported")
	}

	testSuit, err := p.Parse(t.FilePath)
	if err != nil {
		return nil, err
	}

	return []models.TestSuit{testSuit}, nil
}

func (t *ThorrApp) getTestSuitsFromDir() ([]models.TestSuit, error) {
	parser := parser.NewDirParser()

	return parser.Parse(t.FilePath)
}

func (t *ThorrApp) runIntegration() error {
	testSuits, err := t.getTestSuits()
	if err != nil {
		return err
	}

	for _, testSuit := range testSuits {
		fmt.Println(testSuit)
	}

	return nil
}

func (t *ThorrApp) runLoad() error {
	return nil
}
