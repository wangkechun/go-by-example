package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is generated, glhf!")

	fmt.Println("Please input your guess")
	for {
		var input string

		_, err := fmt.Scanf("%s", &input) // fmt.Scanf 输入

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
