package main

import (
	"fmt"
	"math"
)

func main() {

	square := math.Pow(5, 2)
	cube := math.Pow(4, 3)

	fmt.Println("Largest Positive:", math.Max(square, cube))
	fmt.Println("Smallest Positive:", math.Min(square, cube))

	square *= -1
	cube *= -1

	fmt.Println("\nLargest Negative:", math.Max(square, cube))
	fmt.Println("Smallest Negative:", math.Min(square, cube))
}
