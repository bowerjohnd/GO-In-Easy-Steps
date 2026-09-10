package main

import "fmt"

type coords struct {
	x, y int
}

type circle struct {
	radius int
	coords // embedded struct
}

func (c coords) offset() int {
	return c.x*c.x + c.y*c.y
}

func (c circle) diameter() int {
	return c.radius * 2
}

func main() {
	r := circle{radius: 10, coords: coords{5, 5}}

	fmt.Printf("Center: x:%v, y:%v \n", r.x, r.y)
	fmt.Printf("Diameter: %v \n", r.diameter())
	fmt.Printf("Offset from origin: %v \n", r.offset())
}
