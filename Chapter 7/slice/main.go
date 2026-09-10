package main

import "fmt"

func main() {
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	weekend := days[5:]

	fmt.Printf("Slice weekend: %v \n", weekend)
	fmt.Printf("Type weekend: %T \n", weekend)

	fmt.Printf("Length of weekend: %v \n", len(weekend))
	fmt.Printf("Capacity of weekend: %v \n", cap(weekend))

	fmt.Printf("Pointer of weekend: %p \n", weekend)
	fmt.Printf("Address of days[0]: %p \n", &days[5])
}
