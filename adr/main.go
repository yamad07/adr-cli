package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "adr-cli",
		Usage: "write adr faster",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a", "ad"},
				Action:  Add,
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Action:  Init,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
