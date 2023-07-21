package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	var language string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "bwangel"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		// if c.String("lang") == "chinese" {
		if language == "chinese" {
			fmt.Println("你好", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
