package ch6

import "math"

func HeapSort(input []int) (output []int) {
	input = BuildMaxHeap(input)
	for i := len(input) - 1; i >= 1; i-- {
		output = append([]int{input[0]}, output...)
		input = input[1:]
		input = MaxHeap(input, 0)
	}
	return output
}

// BuildMaxHeap creates a max heap complete binary tree
func BuildMaxHeap(input []int) (output []int) {
	halfLength := int(math.Floor(float64(len(input) / 2)))
	for i := halfLength; i >= 0; i-- {
		input = MaxHeap(input, i)
	}
	output = input
	return output
}

// MaxHeap moves the item at the passed in location down to it's appropriate place in the max-heap tree
func MaxHeap(input []int, loc int) (output []int) {
	if loc == 0 {
		MaxHeap(LargestElementFirst(input), 1)
	}
	left := LeftBinaryTreeLocation(loc)
	right := RightBinaryTreeLocation(loc)
	var largest int
	if left <= len(input)-1 && input[left] > input[loc] {
		largest = left
	} else {
		largest = loc
	}
	if right <= len(input)-1 && input[right] > input[largest] {
		largest = right
	}
	if largest != loc {
		input[loc], input[largest] = input[largest], input[loc]
		MaxHeap(input, largest)
	}
	output = input
	return output
}

func LargestElementFirst(input []int) (output []int) {
	maxValLoc := 0
	for i, val := range input {
		if input[maxValLoc] < val {
			maxValLoc = i
		}
	}
	if maxValLoc != 0 {
		input[0], input[maxValLoc] = input[maxValLoc], input[0]
	}
	output = input
	return output
}

func LeftBinaryTreeLocation(loc int) (out int) {
	out = 2 * loc
	return out
}

func RightBinaryTreeLocation(loc int) (out int) {
	out = (2 * loc) + 1
	return out
}
