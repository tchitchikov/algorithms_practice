package mergesort

// DeconstructArray turns [1,2,3] into [[1],[2],[3]]
func DeconstructArray(someArray []int, input [][]int) ([]int, [][]int) {
	if len(input) == 0 {
		input := splitInHalf(someArray)
		return DeconstructArray(someArray, input)
	}
	newInput := [][]int{}
	if len(input) > 1 {
		tempInput := [][]int{}
		for _, array := range input {
			if len(array) > 1 {
				tempInput = splitInHalf(array)
				for _, array := range tempInput {
					newInput = append(newInput, array)
				}
			} else {
				newInput = append(newInput, array)
			}
		}
	}
	if len(newInput) > 0 {
		countLargeArrays := 0
		for _, array2 := range newInput {
			if len(array2) > 1 {
				countLargeArrays = countLargeArrays + 1
			}
		}
		if countLargeArrays > 0 {
			return DeconstructArray(someArray, newInput)
		}
		return someArray, newInput
	}
	return someArray, input
}

func splitInHalf(someArray []int) [][]int {
	output := [][]int{}
	output = append(output, someArray[0:(len(someArray)/2)])
	output = append(output, someArray[(len(someArray)/2):len(someArray)])
	return output
}
