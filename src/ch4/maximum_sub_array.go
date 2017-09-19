package ch4

import (
	"common"
	"fmt"
)

func FindMaxCrossingSubarray(input []int) (maxLeft int, maxRight int, sum int) {
	leftSum := 0
	rightSum := 0
	for iter := common.FindIntArrayMidpoint(input); iter >= 0; iter-- {
		sum = sum + input[iter]
		if sum > leftSum || iter == common.FindIntArrayMidpoint(input) {
			leftSum = sum
			maxLeft = iter
		}
	}
	sum = 0
	for iter := common.FindIntArrayMidpoint(input) + 1; iter < len(input); iter++ {
		sum = sum + input[iter]
		if sum > rightSum || iter == common.FindIntArrayMidpoint(input)+1 {
			rightSum = sum
			maxRight = iter
			fmt.Println(sum, rightSum, maxRight)
		}
	}
	sum = leftSum + rightSum
	return maxLeft, maxRight, sum
}
