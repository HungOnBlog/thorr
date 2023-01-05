package simplecli

import (
	"os"

	"github.com/urfave/cli/v2"
)

type ThorrSimpleCli struct {
	cliApp *cli.App
}

var thorrSimpleCliApp *cli.App

func thorrAction(c *cli.Context) error {
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
