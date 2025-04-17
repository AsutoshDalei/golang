package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args[1:])
	args := os.Args[1:]
	var iargs = []int{}

	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		iargs = append(iargs, val)
	}
	max := 0
	for _, val := range iargs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value: ", max)
}
