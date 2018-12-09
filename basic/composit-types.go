package basic

import "fmt"

func compositeTypes() {
	nums := [3]int{1, 2, 3}
	fmt.Println(nums[1])

	words := []string{`ready`, `steady`}
	words = append(words, `go!`)
	fmt.Println(words[1])

	dict := map[string]int{"one": 1, "two": 2}
	fmt.Println(dict["one"])
}

type person struct {
	name string
	age  int
}

func structs() {
	p := person{name: "Mike", age: 30}
	fmt.Printf("Name: %s, age: %d\n", p.name, p.age)
}
