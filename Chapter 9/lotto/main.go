package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"time"
)

func main() {

	// provide rand with current unix timestamp for unique randomness

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	nums := r.Perm(59)
	var luckyNums [6]int

	for i := 0; i < 6; i++ {
		nums[i]++
		luckyNums[i] = nums[i]
	}

	// nums is integers 1 to 59 in random order
	// sorting nums negates the random order

	// luckyNums is used to pick out first 6 random numbers in nums
	// so that those numbers can be sorted

	str := "Your Six Lucky Numbers: "
	slices.Sort(luckyNums[:])

	for i := 0; i < 6; i++ {
		str += strconv.Itoa(luckyNums[i])

		if i != 5 {
			str += " - "
		}
	}

	fmt.Println("\n", str, "\n")
}
