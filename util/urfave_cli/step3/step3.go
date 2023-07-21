package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "test, t",
			Usage: "上述的test和t会按照 , 解析分开,然后填充反引号里面的内容 `CONTENTIN s`",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
