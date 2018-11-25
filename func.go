package main

import (
	"fmt"
	"strings"
)

func parse(s string) (string, string) {
	parts := strings.Split(s, ",")
	return parts[0], parts[1]
}

func process(data []string, f func(string) (string, string)) {
	for _, s := range data {
		name, address := f(s)
		fmt.Println(name, address)
	}
}

func main() {
	process([]string{"ivan, russia", "mike, usa"}, parse)
	process([]string{"ivan, russia", "mike, usa"}, func(s string) (string, string) {
		return "", ""
	})
}
