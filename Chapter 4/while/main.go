package main

import "fmt"

func main() {

	// equavalent to while loop

	counter := 1
	for counter <= 5 {
		fmt.Println("Iteration", counter)
		counter++
	}

	// equavalent to do while loop

	//i := 5
	i := 0
	for {
		fmt.Println("\t\t\tCountdown", i)
		i--

		if i < 1 {
			fmt.Println("\t\t\t\t\tLift Off!")
			break
		}
	}
}
