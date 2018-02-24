package main

import cli "github.com/urfave/cli"

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "List current versions",
			Action:  current,
		},
	}
}
