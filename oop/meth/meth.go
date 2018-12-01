package meth

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  int
}

func (p person) NickName() string {
	return p.name + strconv.Itoa(p.age)
}

func methods() {
	p := person{name: "Antonio", age: 27}
	fmt.Println(p.NickName())
	fmt.Println(person.NickName(p))
}
