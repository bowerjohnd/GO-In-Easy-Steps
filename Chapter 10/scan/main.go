package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	num := r.Intn(20 + 1)
	guess := 0
	flag := true

	fmt.Print("Guess My Number 1-20: ")

	for flag {
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println(err)
		} else if guess > num {
			fmt.Print("Too High, Try Again: ")
		} else if guess < num {
			fmt.Print("Too Low, Try Again: ")
		} else if guess == num {
			fmt.Println("Correct - My Number Is", num)
			flag = false
		}
	}
}
