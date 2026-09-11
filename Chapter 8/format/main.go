package main

import (
	"fmt"
	"time"
)

func main() {

	dt := time.Now().Round(0)

	fmt.Println("\nDefault Format:", dt)
	fmt.Println("Unix Format:", dt.Format(time.UnixDate))
	fmt.Println("ANSIC Format:", dt.Format(time.ANSIC))
	fmt.Println("RFC3339 Format:", dt.Format(time.RFC3339))
	fmt.Println("Custom Format:", dt.Format("January 2, 2006 [Monday]"))

	fmt.Println("\nUS Format:", dt.Format("January 2, 2006"))
	fmt.Println("UK Format:", dt.Format("2 January, 2006"))

	fmt.Println("\nTime 12-Hour:", dt.Format(time.Kitchen))
	fmt.Println("Time 24-Hour:", dt.Format("15:04"))
}
