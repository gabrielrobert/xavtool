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
			Usage:   "Increment to next version",
			Action:  increment,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type, t",
					Usage: "Increment type (major, minor, patch)",
					Value: "minor",
				},
			},
		},
		{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "Set the current project's version",
			Action:  set,
		},
	}
}
