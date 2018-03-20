package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "xavtool"
	app.Usage = "Command-line utility to automatically increase applications version"
	app.Author = "Gabriel Robert"
	app.Email = "g.robert092@gmail.com"
	app.Version = "1.0.0"
	app.Commands = commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
