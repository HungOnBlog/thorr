package simplecli

import (
	"os"

	"github.com/urfave/cli/v2"
)

type ThorrSimpleCli struct {
	cliApp *cli.App
}

var thorrSimpleCliApp *cli.App

func integrationAction(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}

func loadAction(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}

func thorrAction(c *cli.Context) error {
	command := c.Command.Name

	switch command {
	case "integration":
		return integrationAction(c)
	case "load":
		return loadAction(c)
	default:
		return cli.ShowAppHelp(c)
	}
}

func init() {
	thorrSimpleCliApp = &cli.App{
		Name:      "thorr",
		Usage:     "Thorr <=> no more writing test, integration and load testing tools 🌩️🌩️🌩️🌩️",
		UsageText: "thorr [global options]",
		Commands: []*cli.Command{
			{
				Name:      "integration",
				Usage:     "Run integration test",
				UsageText: "thorr integration [options]",
				Aliases: []string{
					"i",
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "file",
						Usage: "The path of the test suit file, or the directory of the test suit files.",
						Aliases: []string{
							"f",
						},
					},
				},
			},
			{
				Name:      "load",
				Usage:     "Run load test",
				UsageText: "thorr load [options]",
				Aliases: []string{
					"l",
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "file",
						Usage: "The path of the test suit file, or the directory of the test suit files.",
						Aliases: []string{
							"f",
						},
					},
				},
			},
		},
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
