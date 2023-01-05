package simplecli

import (
	"fmt"
	"os"

	"github.com/HungOnBlog/thorr/common/utils"
	"github.com/HungOnBlog/thorr/core/models"
	"github.com/urfave/cli/v2"
)

type ThorrSimpleCli struct {
	cliApp *cli.App
}

var thorrSimpleCliApp *cli.App

func thorrAction(c *cli.Context) error {
	f := c.String("file")

	if f == "" {
		return cli.ShowAppHelp(c)
	}

	isFile := utils.IsFile(f)
	var suits []models.TestSuit
	if isFile {
		content := utils.ReadJsonFile(f)
		suit, err := models.UnmarshalTestSuit(content)
		if err != nil {
			return err
		}

		suits = append(suits, suit)
	} else {
		filePaths := utils.LoadAllFilePathsIn(f)
		for _, filePath := range filePaths {
			content := utils.ReadJsonFile(filePath)
			suit, err := models.UnmarshalTestSuit(content)
			if err != nil {
				return err
			}

			suits = append(suits, suit)
		}
	}

	fmt.Println(suits)
	return nil
}

func init() {
	thorrSimpleCliApp = &cli.App{
		Name:      "thorr",
		Usage:     "Thorr <=> no more writing test, integration and load testing tools 🌩️🌩️🌩️🌩️",
		UsageText: "thorr [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The path of the test suit file, or the directory of the test suit files.",
				Aliases: []string{
					"f",
				},
			},
		},
		Action: thorrAction,
	}
}

func (thorrCli *ThorrSimpleCli) New() ThorrSimpleCli {
	thorrCli.cliApp = thorrSimpleCliApp
	return *thorrCli
}

func (thorrCli *ThorrSimpleCli) Run() error {
	return thorrCli.cliApp.Run(os.Args)
}
