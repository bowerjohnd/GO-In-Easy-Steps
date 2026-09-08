package main

import "fmt"

func main() {
	var zero, num, max int = 0, 0, 1
	var up, dn byte = 'A', 'a'

	fmt.Println("\nzero: ", zero, "\tnum: ", num, "\tmax: ", max)
	fmt.Println("up: ", string(up), "\tdn:", string(dn), "\n")

	result := (num == zero)
	fmt.Println("Equality:\t num == zero\t", result)
	result = (up == dn)
	fmt.Println("Equality:\t up == dn\t", result)

	result = (max != zero)
	fmt.Println("Inequality:\t max != zero\t", result)

	result = (zero > max)
	fmt.Println("Greater:\t zer > max\t", result)
	result = (max <= zero)
	fmt.Println("Less or Equal:\t max <= zero\t", result)

}
