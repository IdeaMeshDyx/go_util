package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "actions for database",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "create a new DB",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Usage: "name of database",
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Printf("db %s and action is %s", c.String("name"), c.Command.Name)
						return nil
					},
				},
				{
					Name:  "drop",
					Usage: "create a new DB",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of database",
						},
					},
					Action: func(c *cli.Context) error {
						fmt.Println(c.Command.Name)
						return nil
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
