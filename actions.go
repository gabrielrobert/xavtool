package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func current(c *cli.Context) error {
	fmt.Printf("current version is 0.0.0")
	return nil
}
