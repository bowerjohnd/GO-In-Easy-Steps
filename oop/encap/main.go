package main

import (
	"fmt"

	"example.com/oop/cube"
)

func main() {
	var box cube.Dims

	box.SetSize(2, 4, 6)

	fmt.Println("Footprint:", box.GetArea())
	fmt.Println("Volume:", box.GetVolume())
}
