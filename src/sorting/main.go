package main

import (
	"fmt"
	"mergesort"
	"shuffle"
)

func main() {
	fmt.Println("Starting sorting")
	a := []int{1, 3, 5, 2, 4, 7, 6, 7, 2, 3, 1, 234, 262}
	b := []int{1, 2, 3, 4, 5}
	output := [][]int{}
	someArr, out := mergesort.DeconstructArray(a, output)
	fmt.Println(someArr, out)
	// fmt.Println(shuffle.ShuffleArray(b))z
	fmt.Println(check_shuffle_distribution(b))
}

func check_shuffle_distribution(input []int) map[int]int {
	distribution := make(map[int]int)
	simulation_length := 10000
	for sim := 0; sim < simulation_length; sim++ {
		first_num := shuffle.ShuffleArray(input)
		distribution[first_num[0]] = distribution[first_num[0]] + 1

	}
	return distribution

}
