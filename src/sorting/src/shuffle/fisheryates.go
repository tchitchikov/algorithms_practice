package shuffle

import (
	"math/rand"
	"time"
)

func ShuffleArray(input []int) []int {
	outputArray := []int{}
	input_length := len(input)
	input_array := make([]int, input_length) // if we don't do this, the input []int is overwritten eventually with all [5,5,5,5,5]'s
	copy(input_array, input)
	for i := 0; i < input_length; i++ {
		randomNum := generateRandom(input_array)
		outputArray = append(outputArray, input_array[randomNum])
		input_array = append(input_array[:randomNum], input_array[(randomNum+1):]...)
	}
	return outputArray
}

func generateRandom(input []int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(len(input))
}
