package wordle

import "fmt"

func ProcessGuess(secretWord string, guess string) ([5]string, bool) {
	if secretWord == guess {
		return [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true
	}
	result := [5]string{}
	for i := range guess {
		if guess[i] == secretWord[i] {
			result[i] = "🟩"
		}
	}

	// yellow
	used := [5]bool{}

	for i := range guess {
		if result[i] != "" {
			// skip over green slots
			continue
		}
		for j := range secretWord {
			if result[j] == "🟩" {
				// skip over filled slots
				continue
			}
			//if !strings.Contains(secretWord, string(guess[i])) {
			//	continue
			//}
			fmt.Println(secretWord, guess, i, j, used)
			if secretWord[j] == guess[i] && !used[j] {
				fmt.Println("Changing to yellow")
				result[i] = "🟨"
				used[j] = true
				//} else {
				//	result[i] = "🟥"
			}
			fmt.Println("  ", used)
		}
	}

	// reds
	for i := range guess {
		if result[i] != "" {
			// skip over filled-in slots
			continue
		}
		result[i] = "🟥"
		//if !strings.Contains(secretWord, string(guess[i])) {
		//	result[i] = "🟥"
		//}
	}

	return result, false
}
