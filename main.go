package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "xavtool"
	app.Usage = "Command-line utility to automatically increase applications version"
	app.Author = "Gabriel Robert"
	app.Email = "g.robert092@gmail.com"
	app.Version = "0.1.0"

	app.Action = func(c *cli.Context) error {
		fmt.Println("cli yay")
		return nil
	}

	app.Run(os.Args)
}
