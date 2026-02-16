package main

import (
	"fmt"
	"math/rand/v2"
)
func numberGenerator() int {
	min := 1
	max := 100
	return (min + rand.IntN(max-min+1))
}
func main(){
	var choice, guess, numGuesses, maxNumGuesses, number int
	var play_again string

	for{
		choice, guess, numGuesses, maxNumGuesses, number = 0, 0, 0, 0, 0
		fmt.Println("Welcome to the number guessing game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		fmt.Println("You have 5 chances to guess the correct number.")

		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("You have selected the Easy difficulty level.")
			maxNumGuesses = 10
		case 2:
			fmt.Println("You have selected the Medium difficulty level.")
			maxNumGuesses = 5
		case 3:
			fmt.Println("You have selected the Hard difficulty! Good luck..")
			maxNumGuesses = 3
		}

		number = numberGenerator()
		for{
			fmt.Print("Enter your guess: ")
			fmt.Scan(&guess)
			numGuesses++
			if guess < number {
    			fmt.Println("Too low! You have", (maxNumGuesses - numGuesses), "guesses left.")
			} else if guess > number {
    			fmt.Println("Too high!You have", (maxNumGuesses - numGuesses), "guesses left.")
			}
			if guess == number{
				fmt.Println("Correct! You win the game!")
				break
			} else if numGuesses == maxNumGuesses{
				fmt.Println("You have run out of guesses!")
				break
			} 
				
		}

		fmt.Println("Do you want to play again? (y/n)")
		fmt.Scan(&play_again)
		if play_again == "y"{
			fmt.Println("Restarting game..")
		}else{
			break
		}
	}

	




}
