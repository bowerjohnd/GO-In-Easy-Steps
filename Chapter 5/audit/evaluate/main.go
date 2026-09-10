package main

import (
	"fmt"

	"example.com/audit/validate"
)

func main() {

	for i := 2; i >= -2; i-- {
		res, err := validate.IsPosInt(i)

		if err != nil {
			fmt.Println("Failed:", err)
		} else {
			fmt.Println(res, "passed")
		}
	}
}
