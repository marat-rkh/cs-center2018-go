package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("./some.txt")
	defer file.Close()
	fmt.Fprintln(file, "data")
}

func createFile(p string) *os.File {
	file, err := os.Create(p)
	if err != nil {
		os.Exit(1)
	}
	return file
}
