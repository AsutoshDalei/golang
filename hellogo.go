package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p

}
func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{age: 20})

	fmt.Println(&person{name: "ann", age: 40})

	fmt.Println(newPerson("Jon"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)
}
