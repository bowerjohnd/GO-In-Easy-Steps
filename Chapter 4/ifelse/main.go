package main

import "fmt"

func main() {

	if 10 > 3 {
		fmt.Println("\nFirst condition is true")
	}

	if 5 > 1 {
		if 7 > 2 {
			fmt.Println("\nBoth nested conditions are ture")
		}
	}

	if 5 < 1 {
		fmt.Println("\nFirst branch is true")
	} else if 2*3 == 6 {
		fmt.Println("\nSecond branch is true")
	} else {
		fmt.Println("\nNeither condition is true")
	}
}
