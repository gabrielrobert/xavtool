package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func saveFile(filePath string, data []byte) {
	os.RemoveAll(filePath)
	err := ioutil.WriteFile(filePath, data, 0666)
	check(err)
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func readFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return dat
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
