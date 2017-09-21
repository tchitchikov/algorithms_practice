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

func FindMaxSubarrayBruteForce(input []int) (output []int, sumValue int, leftPos int, rightPos int) {
	// array size
	for i := 1; i < len(input); i++ {
		// array values
		for j := 0; j+i < len(input); j++ {
			var iteratorArray []int
			iteratorArray = input[j : j+i]
			if sum(iteratorArray) > sumValue {
				sumValue = sum(iteratorArray)
				output = iteratorArray
				leftPos = j
				rightPos = j + i - 1
			}
		}
	}
	return output, sumValue, leftPos, rightPos
}

func MaxSubarrayKadane(input []int) (output []int, sumValue int) {
	var sum int
	for i := 0; i < len(input); i++ {
		if sum+input[i] > 0 {
			sum = sum + input[i]
		} else {
			output = []int{}
			sum = 0
		}
		if sum >= sumValue {
			output = append(output, i)
			sumValue = sum
		}
	}
	if output[0] > 0 && output[len(output)-1]+1 <= len(input)-1 {
		output = input[output[0]-1 : output[len(output)-1]+1]
	} else if output[0] > 0 {
		output = input[output[0]-1 : output[len(output)-1]]
	} else {
		output = input[0:output[len(output)-1]]
	}

	return output, sumValue
}

func sum(input []int) (output int) {
	for _, x := range input {
		output = output + x
	}
	return output
}

func max(input []int) (output int) {
	for _, x := range input {
		if x > output {
			output = x
		}
	}
	return output
}
