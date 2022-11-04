package wordle

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWordleCheckTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wordle Suite")
}

var _ = Describe("Wordle", func() {
	It("Should work", func() {
		Expect("foo").To(Equal("foo"))
	})
})

func TestProcessGuess(t *testing.T) {
	type args struct {
		secretWord string
		guess      string
	}
	tests := []struct {
		name      string
		args      args
		emoji     [5]string
		isCorrect bool
	}{
		{"Correct Guess", args{"HELLO", "HELLO"}, [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true},
		{"One Wrong", args{"PARTS", "DARTS"}, [5]string{"🟥", "🟩", "🟩", "🟩", "🟩"}, false},
		{"One misplaced", args{"HACKS", "SACKS"}, [5]string{"🟨", "🟩", "🟩", "🟩", "🟩"}, false},
		{"One of 2 misplaced", args{"EAGER", "EAGLE"}, [5]string{"🟩", "🟩", "🟩", "🟥", "🟨"}, false},
		{"One right one wrong", args{"EARNS", "EAGLE"}, [5]string{"🟩", "🟩", "🟥", "🟥", "🟥"}, false},
		{"One right one wrong", args{"STILL", "LOLLS"}, [5]string{"🟨", "🟥", "🟥", "🟩", "🟨"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emoji, isCorrect := ProcessGuess(tt.args.secretWord, tt.args.guess)
			if !reflect.DeepEqual(emoji, tt.emoji) {
				t.Errorf("ProcessGuess() emoji = %v, want %v", emoji, tt.emoji)
			}
			if isCorrect != tt.isCorrect {
				t.Errorf("ProcessGuess() isCorrect = %v, want %v", isCorrect, tt.isCorrect)
			}
		})
	}
}
