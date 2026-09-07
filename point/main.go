package main

import "fmt"

func main() {
	num := 20
	ptr := &num
	nptr := new(num)

	fmt.Println("num:\t value:", num, "\taddress:", &num)
	fmt.Println("*ptr:\t value:", *ptr, "\taddress:", &ptr)
	fmt.Println("*nptr:\t value:", *nptr, "\taddress:", nptr)

	*ptr = 100
	*nptr = 200
	fmt.Println("updated... num:", num, " copy:", *nptr)
}
