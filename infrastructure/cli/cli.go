package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/HungOnBlog/thorr/core/thorr"
	"github.com/HungOnBlog/thorr/utils"
	"github.com/urfave/cli/v2"
)

const (
	Logo = `
  __  __               
 / /_/ /  ___  ________
/ __/ _ \/ _ \/ __/ __/
\__/_//_/\___/_/ /_/   
                       
`
)

type ThorCliApp struct {
	cli *cli.App
}

func thorCliAction(c *cli.Context) error {
	command := c.Command.Name
	filePath := c.String("file")
	spawn := c.String("spawn")
	spawnInt := utils.StringToInt(spawn)

	thorrApp := thorr.NewThorrApp(filePath, spawnInt, command)
	fmt.Println(thorrApp)
	err := thorrApp.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewThorCliApp() *ThorCliApp {
	app := &cli.App{
		Name:      Logo,
		Usage:     "Thorr is a load testing tool that helps you to test your API",
		UsageText: "thorr command [global options] [arguments...]",
		Commands: []*cli.Command{
			{
				Name:  "integration",
				Usage: "Run integration test",
				Aliases: []string{
					"i", "it",
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Usage:   "Load tests specs file for folder",
					},
					&cli.StringFlag{
						Name:    "spawn",
						Aliases: []string{"s"},
						Value:   "1",
						Usage:   "Spawn number of goroutines. Default is 1, -1 mean Thorr will spawn as many goroutines as possible",
					},
				},
				Action: thorCliAction,
			},
			{
				Name:  "load",
				Usage: "Run load test",
				Aliases: []string{
					"l",
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Usage:   "Load tests specs file for folder",
					},
					&cli.StringFlag{
						Name:    "spawn",
						Aliases: []string{"s"},
						Value:   "1",
						Usage:   "Spawn number of goroutines. Default is 1, -1 mean Thorr will spawn as many goroutines as possible",
					},
				},
				Action: thorCliAction,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Load tests specs file for folder",
			},
			&cli.StringFlag{
				Name:    "spawn",
				Aliases: []string{"s"},
				Value:   "1",
				Usage:   "Spawn number of goroutines. Default is 1, -1 mean Thorr will spawn as many goroutines as possible",
			},
		},
		Action: thorCliAction,
	}

	return &ThorCliApp{
		cli: app,
	}
}

func (t *ThorCliApp) Run() error {
	if err := t.cli.Run(os.Args); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
