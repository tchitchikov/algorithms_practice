package main

import (
	"ch4"
	"ch6"
	"common"
	"fmt"
)

func main() {
	inputArray := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}
	// inputArray := []int{100, 110, 111, 99, 98, 97}
	// inputArray := []int{6, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	// fmt.Println(ch4.FindMaxCrossingSubarray(common.TimeSeriesMove(inputArray), 0, common.FindIntArrayMidpoint(inputArray), len(inputArray)))
	// fmt.Println(ch4.FindMaxSubarray(common.TimeSeriesMove(inputArray), 0, len(common.TimeSeriesMove(inputArray))-1))
	fmt.Println(ch4.FindMaxSubarrayBruteForce(common.TimeSeriesMove(inputArray)))
	fmt.Println(ch4.MaxSubarrayKadane(common.TimeSeriesMove(inputArray)))
	fmt.Println(ch6.HeapSort(inputArray))
}
