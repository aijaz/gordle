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

	//
	//index := rand.Intn(len(validWords))
	//for k := range validWords {
	//	if index == 0 {
	//		return k
	//	}
	//	index--
	//}
	//return ""
}

func GetGuess() (string, error) {
	fmt.Println("Guess a word: ")
	var word string
	_, err := fmt.Scan(&word)
	if err != nil {
		return "", err
	}
	word = strings.ToUpper(word)
	if !validWords[word] {
		return "", errors.New("this is not a valid word")
	}
	if len(word) != 5 {
		return "", errors.New("all guesses should be 5 characters in length")
	}

	return word, nil
}

func populateWords() {
	validWords["ARISE"] = true
	validWords["BROOK"] = true
	validWords["COUNT"] = true
	validWords["DENIM"] = true
	validWords["ENJOY"] = true
	validWords["FLOAT"] = true
	validWords["OCEAN"] = true

	file, err := os.Open("wordle/all5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, ",")
	for _, word := range words {
		validWords[word] = true
	}
}
