package main

import "fmt"

func main() {
	var yes, aye, no bool = true, true, false

	result := (yes && aye)
	fmt.Println("AND Logic:\tyes && aye\t", result)
	result = (yes && no)
	fmt.Println("AND Logc:\tyes && no\t", result)
	result = (yes || no)
	fmt.Println("OR Logic:\tyes || no\t", result)

	result = !yes
	fmt.Println("NOT Logic:\t yes =", yes, "\t!yes = ", result)

}
