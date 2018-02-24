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
		{
			Name:    "increment",
			Aliases: []string{"i"},
			Usage:   "Increment to next minor version",
			Action:  increment,
		},
	}
}
