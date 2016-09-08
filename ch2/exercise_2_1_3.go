package main

import (
	"fmt"
)

func main() {
	A := []string{"matrix", "matricies", "linear transformations", "vector"}
	fmt.Println(LinearSearch(A, "matrix"))
	fmt.Println(LinearSearch(A, "coffee"))
}

func LinearSearch (A []string, term string) int {
	for i, val := range A {
		if term == val {
			return i
		}
	}
	return -1 
}
