package sort

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type personSlice []person

func (ps personSlice) Len() int           { return len(ps) }
func (ps personSlice) Less(i, j int) bool { return ps[i].age < ps[j].age }
func (ps personSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	ps := personSlice{
		person{name: "Bogdan", age: 26},
		person{name: "Antonio", age: 27},
		person{name: "Marat", age: 21},
	}
	sort.Sort(ps)
	fmt.Println(ps) // Prints: [{Marat 21} {Bogdan 26} {Antonio 27}]
}
