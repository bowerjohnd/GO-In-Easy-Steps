package main

import (
	"flag"
	"fmt"
)

func main() {

	txt := flag.String("txt", "C#", "A string")
	num := flag.Int("num", 8, "An integer")
	sta := flag.Bool("sta", false, "A Boolean")

	flag.Parse()
	fmt.Println("Text:", *txt)
	fmt.Println("Number:", *num)
	fmt.Println("Status:", *sta)

}
