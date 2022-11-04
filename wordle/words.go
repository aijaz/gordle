package wordle

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var validWords = make(map[string]bool)

func GetSecretWord() string {
	if len(validWords) == 0 {
		populateWords()
	}

	fileBytes, err := os.ReadFile("wordle/secretWords.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	secretWords := strings.Split(string(fileBytes), "\n")
	index := rand.Intn(len(secretWords))
	return strings.ToUpper(secretWords[index])
}

func GetGuess() (string, error) {
	fmt.Println("Guess a word: ")
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		return "", err
	}
	word = strings.ToUpper(word)
	if len(word) != 5 {
		return "", errors.New("all guesses should be 5 characters in length")
	}
	if !validWords[word] {
		return "", errors.New("this is not a valid word")
	}

	return word, nil
}

func populateWords() {
	file, err := os.Open("wordle/all5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, ",")
	for _, word := range words {
		validWords[word] = true
	}
}
