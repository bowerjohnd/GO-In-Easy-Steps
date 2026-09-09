package main

import "fmt"

func iterate(num int) {

	if num < 1 {
		fmt.Println("\t\t\t\tLift Off!")
	} else {
		fmt.Println("\t\tCountdown", num)
		num--
		iterate(num)
	}
}

func main() {
	iterate(5)
}
