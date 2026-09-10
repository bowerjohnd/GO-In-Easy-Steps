package main

import "fmt"

func main() {

	arr := [...]int{100, 200, 300, 400, 500}

	fmt.Println("No. of elements:", len(arr))

	for i, v := range arr {
		arr[i] = v / 10
	}

	for i, v := range arr {
		fmt.Printf("Index: %v Value: %v \n", i, v)
	}

	for _, v := range arr { // when you don't need the index number
		fmt.Println("Blank Identifier - value:", v)
	}
}
