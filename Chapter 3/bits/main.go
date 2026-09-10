package main

import "fmt"

func main() {
	var flags byte = 10

	fmt.Printf("Flag #1: %v \n", (flags&1) > 0)
	fmt.Printf("Flag #2: %v \n", (flags&2) > 0)
	fmt.Printf("Flag #3: %v \n", (flags&4) > 0)
	fmt.Printf("Flag #4: %v \n", (flags&8) > 0)

	fmt.Printf("Flags Value:\t %08b \t %v \n", flags, flags)
	flags = flags >> 1
	fmt.Printf("Shifted Right:\t %08b \t %v \n", flags, flags)

}
