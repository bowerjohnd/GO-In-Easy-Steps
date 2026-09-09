package main

import "fmt"

func cube(num int) (string, int, int) {
	return "Result", num, (num * num * num)
}

func main() {
	//a, b, c := cube(5)
	_, b, c := cube(5)
	fmt.Println(b, "Cubed =", c)
}
