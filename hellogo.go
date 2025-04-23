package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str1 string
}

func main() {
	co := container{
		base: base{num: 1},
		str1: "some name",
	}

	b := base{num: 2}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str1)

	fmt.Println("also num: ", co.base.num)

	fmt.Println("Describe 1: ", co.describe())
	fmt.Println("Describe 2: ", b.describe())
	type describer interface {
		describe() string
	}

	var d describer = co

	fmt.Println("Describe: ", d.describe())

}
