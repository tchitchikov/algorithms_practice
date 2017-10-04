package main

import "fmt"

func main() {
	// how do you get the memory address of a variable
	x := 5
	fmt.Println(&x)

	// how do you assign a value to a pointer
	fmt.Println(x)
	fmt.Println(&x)
	square(&x)
	fmt.Println(&x)
	fmt.Println(x)
	zero_out(&x)
	fmt.Println(x)

	// create a new pointer
	y := new(int)
	fmt.Println(y)
	fmt.Println(*y)
	*y = 123

	fmt.Println(x, *y)
	swap(&x, y)
	fmt.Println(x, *y)

}

func zero_out(input *int) {
	*input = 0
}

func square(x *int) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	z := new(int)
	*z = *x
	*x = *y
	*y = *z
}
