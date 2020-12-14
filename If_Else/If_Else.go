package main

import "fmt"

func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}

func main() {
	number := 50
	guess := -5

	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The Guess must be between 1 and 100")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("To Low")
		}
		if guess > number {
			fmt.Println("To High")
		}
		if guess == number {
			fmt.Println("Equal")
		}

	}

}
