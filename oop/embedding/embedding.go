package embedding

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

type employee struct {
	person
	jobTitle string
}

func methods() {
	emp := employee{person: person{name: "Antonio", age: 27}, jobTitle: "Coder"}
	fmt.Println(emp.person.name, emp.person.age, emp.jobTitle)
	fmt.Println(emp.name, emp.age, emp.jobTitle)
	fmt.Println(emp.NickName())
	var _ person = employee{}
}
