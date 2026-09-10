package main

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Yellow = "\033[33m"
)

func main() {

	colors := make(map[string]string)

	colors["Red"] = "#FF0000"
	colors["Green"] = "#00FF00"
	colors["Blue"] = "#0000FF"

	fmt.Printf("\nColors: %v \n", colors)
	fmt.Println()

	colors["Yellow"] = "#FFFF00"
	fmt.Printf("%sYellow%s Hex Code: %v\n", Yellow, Reset, colors["Yellow"])
	fmt.Println()

	delete(colors, "Blue")

	fontColor := Reset // added for fun

	// iterates through keys in random order
	for k, v := range colors {

		// added for fun
		switch k {
		case "Red":
			fontColor = Red
		case "Green":
			fontColor = Green
		case "Blue":
			fontColor = Blue
		case "Yellow":
			fontColor = Yellow
		}

		fmt.Printf("Hex Code for %s %v %s is %v \n", fontColor, k, Reset, v)
	}

	fmt.Println()
	colors["Blue"] = "#0000FF"

	// iterates through keys in random order
	for k, v := range colors {

		// added for fun
		switch k {
		case "Red":
			fontColor = Red
		case "Green":
			fontColor = Green
		case "Blue":
			fontColor = Blue
		case "Yellow":
			fontColor = Yellow
		}

		fmt.Printf("Hex Code for %s %v %s is %v \n", fontColor, k, Reset, v)
	}
}
