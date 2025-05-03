package main

import (
	"encoding/json"
	"fmt"
)

type resp1 struct {
	Page   int
	Fruits []string
}

type resp2 struct {
	Page   int
	Fruits []string
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	bolB, _ = json.Marshal("hello")
	fmt.Println(string(bolB))

	strB := []string{"apple", "pear", "mango"}
	bolB, _ = json.Marshal(strB)
	fmt.Println(string(bolB))

	strC := map[string]int{"apple": 5, "mango": 10}
	bolB, _ = json.Marshal(strC)
	fmt.Println(string(bolB))

	res1D := &resp1{
		Page:   1,
		Fruits: []string{"apple", "pear", "banana"},
	}
	res1b, _ := json.Marshal(res1D)
	fmt.Println(string(res1b))

}
