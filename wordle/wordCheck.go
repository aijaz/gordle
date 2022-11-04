package wordle

import "strings"

func ProcessGuess(secretWord string, guess string) ([5]string, bool) {
	if secretWord == guess {
		return [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true
	}
	result := [5]string{}
	for i, _ := range guess {
		if guess[i] == secretWord[i] {
			result[i] = "🟩"
		} else if strings.Contains(secretWord, string(guess[i])) {
			result[i] = "🟨"
		} else {
			result[i] = "🟥"
		}
	}
	return result, false
}
