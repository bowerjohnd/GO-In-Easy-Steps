package main

import "fmt"

type car struct {
	color string
	body  string
}

func (c car) motion(stop bool) string {
	if stop {
		return "<--- stopping"
	}
	return "accelerating --->"
}

func main() {
	porsche := car{color: "blue", body: "coupe"}
	bentley := car{color: "green", body: "saloon"}

	fmt.Println("Porsche paint is", porsche.color)
	fmt.Println("Porsche style is", porsche.body)
	fmt.Println("Porsche is", porsche.motion(false))

	fmt.Println("Bentley paint is", bentley.color)
	fmt.Println("Bentley style is", bentley.body)
	fmt.Println("Bentley is", bentley.motion(true))
}
