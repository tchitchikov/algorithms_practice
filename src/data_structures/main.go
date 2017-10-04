package main

import (
	"fmt"
	"hashtablitza"
)

func main() {
	someArray := [5]string{"George", "John", "Thomas", "James", "Andrew"}
	val1 := hashtablitza.HashIndex(someArray, "John")
	fmt.Println(val1)
}
