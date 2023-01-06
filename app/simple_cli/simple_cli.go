package simplecli

import (
	"fmt"
	"os"

	"github.com/HungOnBlog/thorr/core/models"
	"github.com/urfave/cli/v2"
)

const LOGO = `

___________.__                         
\__    ___/|  |__   __________________ 
  |    |   |  |  \ /  _ \_  __ \_  __ \
  |    |   |   Y  (  <_> )  | \/|  | \/
  |____|   |___|  /\____/|__|   |__|   
                \/                     
`

type ThorrSimpleCli struct {
	cliApp *cli.App
}

var thorrSimpleCliApp *cli.App

func integrationAction(c *cli.Context) error {
	file := c.String("file")
	spawns := c.Int("spawns")

	options := &models.ThorrOptions{
		File:    file,
		Spawns:  spawns,
		Command: "integration",
	}

	thorr := models.NewThorr(*options)
	err := thorr.Start()
	if err != nil {
		return err
	}

	return nil
}

func loadAction(c *cli.Context) error {
	return nil
}

func thorrAction(c *cli.Context) error {
	command := c.Command.Name
	fmt.Println("COMMAND: ", command)

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
		Name:      LOGO,
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
						Value: "",
						Aliases: []string{
							"f",
						},
					},
					&cli.IntFlag{
						Name:  "spawns",
						Usage: "The number of spawns. Default 1, max 1000. Special value -1 Thorr will spawn the routine number equal to the number of test suits.",
						Value: 1,
						Aliases: []string{
							"s",
						},
					},
				},
				Action: thorrAction,
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
						Value: "",
						Aliases: []string{
							"f",
						},
					},
					&cli.IntFlag{
						Name:  "spawns",
						Usage: "The number of spawns. Default 1, max 1000. Special value -1 Thorr will spawn the routine number equal to the number of test suits.",
						Value: 1,
						Aliases: []string{
							"s",
						},
					},
				},
				Action: thorrAction,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Usage: "The path of the test suit file, or the directory of the test suit files.",
				Value: "",
				Aliases: []string{
					"f",
				},
			},
			&cli.IntFlag{
				Name:  "spawns",
				Usage: "The number of spawns. Default 1, max 1000. Special value -1 Thorr will spawn the routine number equal to the number of test suits.",
				Value: 1,
				Aliases: []string{
					"s",
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
