package main

import "fmt"

func main() {
	num := 2
	char := 'B'

	switch num {
	case 1:
		fmt.Println("\nNumber is One")
	case 2:
		fmt.Println("\nNumber is Two")
	case 3:
		fmt.Println("\nNumber is Three")
	default:
		fmt.Println("\nNumber is Unrecognized")
	}

	switch char {
	case 'A':
		fmt.Println("\nLetter is A")
	case 'B':
		fmt.Println("\nLetter is B")
	default:
		fmt.Println("\nLetter is Unrecognized")
	}

	// added for fun

	googleExampleOfFallthroughUseSwitch(2)
	googleExampleOfFallthroughUseSwitch(3)
	googleExampleOfFallthroughUseSwitch(1)
	googleExampleOfFallthroughUseSwitch(90)

}

func googleExampleOfFallthroughUseSwitch(dangerLevel int) {

	fmt.Printf("\nHandling Danger Level %d:\n", dangerLevel)

	switch dangerLevel {
	case 3:
		fmt.Println("\t[-] Evacuate the facility!")
		fallthrough
	case 2:
		fmt.Println("\t[-] Disconnect External Servers.")
		fallthrough
	case 1:
		fmt.Println("\t[-] Log the system anomaly to admin.")
	default:
		fmt.Println("\t[-] Monitoring normal operations.")
	}
}
