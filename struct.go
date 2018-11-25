package main

import "fmt"

type person struct {
	name string
	age  int
}

func structs() {
	p := person{
		name: "Mike",
		age:  30,
	}
	fmt.Printf("Name: %s, age: %d\n", p.name, p.age)
}
