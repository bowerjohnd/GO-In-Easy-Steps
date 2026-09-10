package main

import "fmt"

func main() {
	const Pi = 3.14159

	const (
		Red = iota + 1
		Yellow
		Green
		Brown
		Blue
		Pick
		Black
	)

	fmt.Printf("Pi approximately: %v \n\n", Pi)
	fmt.Printf("Red: %v point \n", Red)
	fmt.Printf("Blue: %v points \n", Blue)
	fmt.Printf("Black: %v points \n", Black)
}
