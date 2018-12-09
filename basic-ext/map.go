package basic_ext

import "strings"

func maps() {
	m := map[string]int{}
	m["One"] = 1
	m["two"] = 2

	m2 := map[string]int{}
	m2[key("One")] = 1
	m2[key("two")] = 2
}

func key(s string) string {
	return strings.ToLower(s)
}
