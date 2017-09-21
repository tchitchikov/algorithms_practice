package main

import (
	"ch4"
	"common"
	"fmt"
)

func main() {
	inputArray := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}
	// inputArray := []int{100, 110, 111, 99, 98, 97}
	// fmt.Println(ch4.FindMaxCrossingSubarray(common.TimeSeriesMove(inputArray), 0, common.FindIntArrayMidpoint(inputArray), len(inputArray)))
	// fmt.Println(ch4.FindMaxSubarray(common.TimeSeriesMove(inputArray), 0, len(common.TimeSeriesMove(inputArray))-1))
	fmt.Println(ch4.FindMaxSubarrayBruteForce(common.TimeSeriesMove(inputArray)))
}
