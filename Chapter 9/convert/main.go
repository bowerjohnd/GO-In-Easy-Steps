package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "I have Nothing to declare except My Genius"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	str = " 42 "
	fmt.Printf("\n'%v' Type: %T, Length: %v \n", str, str, len(str))
	str = strings.Trim(str, " ")
	fmt.Printf("\n'%v' Type: %T, Length: %v \n", str, str, len(str))

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Atoi: %v Type: %T \n", num, num)
		num := strconv.Itoa(num)
		fmt.Printf("Itoa: %v Type: %T \n", num, num)
	}

}
