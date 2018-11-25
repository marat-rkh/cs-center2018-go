package main

import (
	"fmt"
	"os"
)

func writeFile(data []string) {
	file, err := os.Create("./some.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close() // !!!
	if len(data) == 0 {
		fmt.Println("nothing to write")
		return
	}
	for _, s := range data {
		fmt.Fprintln(file, s)
	}
}
