package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yamad07/adr-cli"
)

func Add(c *cli.Context) error {
	name := c.Args().Get(0)
	if err := adr.CreateNewRecord(name); err != nil {
		return err
	}

	return nil
}

func Init(c *cli.Context) error {
	root := c.Args().Get(0)
	if err := adr.CreateConfigFile(root); err != nil {
		return err
	}

	return nil
}
