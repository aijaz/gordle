package wordle

import (
	"strings"
)

func ProcessGuess(secretWord string, guess string) ([5]string, bool) {
	if secretWord == guess {
		return [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true
	}
	result := [5]string{}
	available := make(map[string]int)
	for c := range secretWord {
		available[string(secretWord[c])]++
	}

	for i := range guess {
		currentCharacter := string(guess[i])
		if guess[i] == secretWord[i] {
			result[i] = "🟩"
			available[currentCharacter]--
		}
	}
	for i := range guess {
		if result[i] == "🟩" {
			continue
		}
		currentCharacter := string(guess[i])
		if strings.Contains(secretWord, currentCharacter) && available[currentCharacter] > 0 {
			result[i] = "🟨"
			available[currentCharacter]--
		} else {
			result[i] = "🟥"
		}
	}

	return result, false
}
