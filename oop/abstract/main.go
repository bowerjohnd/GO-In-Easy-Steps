package main

import "fmt"

type bird interface {
	speak() string
	move() string
}

type human interface {
	speak() string
	move() string
}

type creature interface {
	bird
	human
}

type parrot struct{}

func (parrot) speak() string {
	return "Squawk, squawk!"
}
func (parrot) move() string {
	return "A parrot flies away."
}

type person struct {
	hair string
	eyes string
}

func (person) speak() string {
	return "Hi, there!"
}
func (person) move() string {
	return "A person walks away."
}

func nudge(c creature) {
	fmt.Printf("\n%v \n", c.speak())
	fmt.Printf("%v \n\n", c.move())
}

func main() {
	var bird1 parrot
	var human1 person

	nudge(bird1)
	nudge(human1)

	human1.hair = "blonde"
	human1.eyes = "green"
	fmt.Println("\nHuman1 hair color:", human1.hair)
	fmt.Println("Human1 eye color:", human1.eyes)
}
