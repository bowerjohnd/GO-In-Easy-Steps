package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//txt, err := ioutil.ReadFile("Oscar.txt")	// ioutil deprecated
	txt, err := os.ReadFile("Oscar.txt")
	check(err)
	fmt.Println(string(txt))

	file, err := os.Open("Oscar.txt")
	check(err)
	defer file.Close()

	pos, err := file.Seek(41, 0)
	check(err)

	slice := make([]byte, 15)
	nb, err := file.Read(slice)
	check(err)

	fmt.Printf("\n%v byes @ %v: ", nb, pos)
	fmt.Printf("%v\n", string(slice[:nb]))

}
