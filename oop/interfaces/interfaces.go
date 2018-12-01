package interfaces

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

type NickNamed interface {
	NickName() string
}

func joinChat(nn NickNamed) {
	fmt.Println(nn.NickName() + " has joined")
}

func interfaces() {
	p := person{name: "Bogdan", age: 26}
	joinChat(p)
}
