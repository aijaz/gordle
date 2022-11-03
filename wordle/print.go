package wordle

import (
	"fmt"
	"strings"
)

func PrintBoard(emoji [6][5]string, words [6]string, guessNumber int) {
	for i := 0; i < guessNumber; i++ {
		fmt.Println(strings.ToUpper(words[i]), strings.Join(emoji[i][:], " "))
	}
}
