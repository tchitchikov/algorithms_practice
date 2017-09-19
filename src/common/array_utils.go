package common

import (
	"math"
)

func FindIntArrayMidpoint(input []int) (midPoint int) {
	return int(math.Floor(float64(len(input)) / 2.0))
}

func TimeSeriesMove(input []int) (output []int) {
	for i := 0; i < len(input)-1; i++ {
		output = append(output, input[i+1]-input[i])
	}
	return output
}
