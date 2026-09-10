package main

import "fmt"

func main() {
	sum := 2*3 + 4 - 5
	fmt.Printf("Default Order: %v", sum)

	sum = 2 * ((3 + 4) - 5)
	fmt.Printf("\n\tForced Order: %v", sum)

	sum = 7 % 3 * 2
	fmt.Printf("\n\nDefault Order: %v", sum)
	sum = 7 % (3 * 2)
	fmt.Printf("\n\tForced Order: %v", sum)
}
