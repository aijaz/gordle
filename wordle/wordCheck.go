package wordle

import "strings"

func ProcessGuess(secretWord string, guess string) ([5]string, bool) {
	if secretWord == guess {
		return [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true
	}
	result := [5]string{}
	for i := range guess {
		if guess[i] == secretWord[i] {
			result[i] = "🟩"
		} else if !strings.Contains(secretWord, string(guess[i])) {
			result[i] = "🟥"
		}
	}
	for i := range guess {
		if result[i] != "" {
			// skip over filled-in slots
			continue
		}
		for j := range secretWord {
			if result[j] != "" {
				continue
			}
			if secretWord[j] == guess[i] {
				result[i] = "🟨"
			} else {
				result[i] = "🟥"
			}
		}
	}
	return result, false
}
