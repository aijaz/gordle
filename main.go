package main

import (
	"fmt"
	"wordleTest/wordle"
)

func main() {
	guessNumber := 0
	guessedCorrectly := false
	secretWord := wordle.GetSecretWord()
	var guess string
	var err error
	var historyEmoji [6][5]string
	var historyString [6]string

	for {
		var emoji [5]string

		guess, err = wordle.GetGuess()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		emoji, guessedCorrectly = wordle.ProcessGuess(secretWord, guess)
		historyEmoji[guessNumber] = emoji
		historyString[guessNumber] = guess

		guessNumber += 1
		if guessNumber == 6 || guessedCorrectly {
			break
		}
		wordle.PrintBoard(historyEmoji, historyString, guessNumber)
	}
	wordle.PrintBoard(historyEmoji, historyString, guessNumber)
	if guessedCorrectly {
		fmt.Println("You win!")
	} else {
		fmt.Println("Better luck next time")
	}
}
