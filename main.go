package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	random := randomNumber()

	for matching := false; !matching; {
		player := userInput()

		matching = match(random, player)

		fmt.Println(matching)
	}
}

func userInput() int {
	var input int

	fmt.Println("Hey, input your guess between 0 & 10:")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Try again ...")
	} else {
		fmt.Println("Your guess:")
	}
	return input

}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(11)
}

func match(random, player int) bool {
	if random > player {
		fmt.Println("Your score is smaller")
		return false
	} else if random > player {
		fmt.Println("Your score is bigger")
		return false
	} else {
		fmt.Println("YOU GOT IT!!!!")
		return true
	}
}
