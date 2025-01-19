package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNum := rand.Intn(100)

	for {
		var input int
		fmt.Println("Please input your guess")
		fmt.Scanf("%v", &input)
		if input > secretNum {
			fmt.Println("Please lower your guess")
		} else if input < secretNum {
			fmt.Println("Please upper your guess")
		} else if input == secretNum {
			fmt.Println("Congratrulations")
			break
		}

	}
}
