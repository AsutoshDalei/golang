package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println(sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Printf("%s -> %s\n", k, "v")
	}

	for i, c := range "google" {
		fmt.Println(i, c)
	}

}
