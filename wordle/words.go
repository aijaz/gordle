package wordle

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var words = make(map[string]bool)

func GetSecretWord() string {
	if len(words) == 0 {
		populateWords()
	}
	index := rand.Intn(len(words))
	for k := range words {
		if index == 0 {
			return k
		}
		index--
	}
	return ""
}

func GetGuess() (string, error) {
	fmt.Println("Guess a word: ")
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		return "", err
	}
	word = strings.ToUpper(word)
	if !words[word] {
		return "", errors.New("this is not a valid word")
	}
	if len(word) != 5 {
		return "", errors.New("all guesses should be 5 characters in length")
	}

	return word, nil
}

func populateWords() {
	words["ARISE"] = true
	words["BROOK"] = true
	words["COUNT"] = true
	words["DENIM"] = true
	words["ENJOY"] = true
	words["FLOAT"] = true
	words["OCEAN"] = true
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
