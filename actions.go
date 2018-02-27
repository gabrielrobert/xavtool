package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func current(c *cli.Context) error {
	dir := getWorkingDir()
	allFiles := findManifests(dir)
	if len(allFiles) == 0 {
		fmt.Println("No application has been found")
	} else {
		for _, file := range allFiles {
			fmt.Println(fmt.Sprintf("%v - %v (%v)", file.Version, file.Name, file.Path))
		}
	}

	return nil
}

func increment(c *cli.Context) error {
	dir := getWorkingDir()
	allFiles := findManifests(dir)
	for _, file := range allFiles {
		newVersion := incrementPatch(file.Version)
		fmt.Println(fmt.Sprintf("%v: New version: %v (%v)", file.Version, newVersion, file.Path))
	}
	return nil
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
