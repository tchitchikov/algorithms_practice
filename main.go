package main

import (
	"ch4"
	"common"
	"fmt"
)

func main() {
	fmt.Println("here")
	inputArray := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}
	fmt.Println(ch4.FindMaxCrossingSubarray(common.TimeSeriesMove(inputArray)))
}
