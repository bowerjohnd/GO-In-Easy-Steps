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

	txt := []byte("\nA thousand suns will stream on thee," +
		"\nA thousand moons will quiver.\n")
	err := os.WriteFile("Farewell.txt", txt, 0644) // book uses deprecated ioutil
	check(err)

	file, err := os.OpenFile("Farewell.txt", os.O_APPEND, 0644)
	check(err)
	defer file.Close()

	slice := []byte("by Alfred Lord Tennyson.\n")
	nb, err := file.Write(slice)
	check(err)

	fmt.Printf("\nAppended: %v byes - %v", nb, string(slice[:nb]))
}
