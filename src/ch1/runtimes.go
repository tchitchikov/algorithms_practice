package main

import (
	"fmt"
	"math"
)

func main() {
	x := [3]float64{3.0, 3600.0, 86400.0}
	for _, n := range x {
		fmt.Printf("%f seconds\n", n)
		fmt.Printf("log(n): %f\n", math.Log10(n))
		fmt.Printf("Sqrt(n): %f\n", math.Sqrt(n))
		fmt.Printf("n: %f\n", n)
		fmt.Printf("n * log(n): %f\n", n*math.Log10(n))
		fmt.Printf("n^2: %f\n", math.Pow(n, 2.0))
		fmt.Printf("n^3: %f\n", math.Pow(n, 3.0))
		fmt.Printf("2^n: %f\n", math.Pow(2.0, n))
		fmt.Printf("n!: %f\n\n\n", Factorial(n))
	}
}

func Factorial(x float64) (z float64) {
	for y := x; y >= 0; y-- {
		if z == 0 {
			z = y
		} else if y > 0 {
			z = z * y
		} else {
			return z
		}
	}
	return 0
}
