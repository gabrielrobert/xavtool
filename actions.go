package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func current(c *cli.Context) error {
	fmt.Printf("current version is %v", getCurrentVersion())
	return nil
}

func getCurrentVersion() string {
	return "1.0.1"
}
