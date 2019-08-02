package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Falcon Lib"
	app.Usage = "Run Falcon Lib"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l, la",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name: "symbol, s, sym",
			Value: "***",
			Usage: "Prefix",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Subcommands: []cli.Command{
				{
					Name:  "new",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name: "old",
			Aliases: []string{"o", "ol"},
			Subcommands: []cli.Command {
				{
					Name: "what",
					Action: func(c *cli.Context) error {
						fmt.Println("What in old: ", c.Args().First())
						return nil
					},
				},
				{
					Name: "how",
					Action: func(c *cli.Context) error {
						fmt.Println("How in old: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name, c.String("symbol"))
		} else {
			fmt.Println("Hello", name, c.String("symbol"))
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}