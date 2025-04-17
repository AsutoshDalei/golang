package main

import (
	"fmt"
)

const pi = 3.14

func main() {
	i := 5
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("beyond")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("idk", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("asxa")

}
