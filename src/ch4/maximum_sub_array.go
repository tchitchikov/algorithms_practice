package ch4

import (
	"math"
)

func FindMaxCrossingSubarray(input []int, low int, mid int, high int) (maxLeft int, maxRight int, sum int) {
	leftSum := 0
	rightSum := 0
	for iter := mid; iter >= low; iter-- {
		sum = sum + input[iter]
		if sum > leftSum || iter == mid {
			leftSum = sum
			maxLeft = iter
		}
	}
	sum = 0
	for iter := mid + 1; iter < high-1; iter++ {
		sum = sum + input[iter]
		if sum > rightSum || iter == mid+1 {
			rightSum = sum
			maxRight = iter
		}
	}
	sum = leftSum + rightSum
	return maxLeft, maxRight, sum
}

func FindMaxSubarray(input []int, low int, high int) (int, int, int) {
	if high == low {
		return low, high, input[low]
	}
	mid := int(math.Floor(float64((low + high) / 2)))
	leftLow, leftHigh, leftSum := FindMaxSubarray(input, low, mid)
	rightLow, rightHigh, rightSum := FindMaxSubarray(input, mid+1, high)
	crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(input, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	}
	return crossLow, crossHigh, crossSum
}
