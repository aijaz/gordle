package wordle

func ProcessGuess(secretWord string, guess string) ([5]string, bool) {
	if secretWord == guess {
		return [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true // 🟥🟨
	}
	return [5]string{"🟥", "🟥", "🟥", "🟥", "🟥"}, false
}

func isGuessLengthValid(guess string) bool {
	return len(guess) == 5
}
