package main

import (
	"testing"
)

func Test_openFile(t *testing.T) {
	filepath := "test/Info.plist"
	file := openFile(filepath)
	if file == nil {
		t.Errorf("file %v should be able to be open", filepath)
	}

	stats, error := file.Stat()
	if error != nil {
		t.Errorf("file %v statistics should be available", filepath)
	}

	size := stats.Size()
	if size <= 0 {
		t.Errorf("file %v should not be empty", filepath)
	}
}

func Test_readFile(t *testing.T) {
	filepath := "test/Info.plist"
	file := readFile(filepath)
	if file == nil {
		t.Errorf("file %v should be able to be read", filepath)
	}
	if len(file) <= 0 {
		t.Errorf("file %v not be empty", filepath)
	}
}
