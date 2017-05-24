package main

import (
	"fmt"
	"mergesort"
)

func main() {
	fmt.Println("Starting sorting")
	a := []int{1,3,5,2,4,7,6,7,2,3,1,234,262}
	output := [][]int{}
	someArr, out := mergesort.DeconstructArray(a, output)
	fmt.Println(someArr, out)
}