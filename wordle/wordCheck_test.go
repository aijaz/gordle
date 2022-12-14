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

var _ = DescribeTable("Running ProcessGuess with various inputs",
	func(secretWord string, guess string, emoji [5]string, isCorrect bool) {
		em, isCo := ProcessGuess(secretWord, guess)
		Expect(em).To(Equal(emoji))
		Expect(isCo).To(Equal(isCorrect))
	},
	Entry("HELLO HELLO", "HELLO", "HELLO", [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true),
	Entry("PARTS DARTS", "PARTS", "DARTS", [5]string{"🟥", "🟩", "🟩", "🟩", "🟩"}, false),
	Entry("HACKS SACKS", "HACKS", "SACKS", [5]string{"🟥", "🟩", "🟩", "🟩", "🟩"}, false),
	Entry("EAGER EAGLE", "EAGER", "EAGLE", [5]string{"🟩", "🟩", "🟩", "🟥", "🟨"}, false),
	Entry("EARNS EAGLE", "EARNS", "EAGLE", [5]string{"🟩", "🟩", "🟥", "🟥", "🟥"}, false),
	Entry("STILL LOLLS", "STILL", "LOLLS", [5]string{"🟨", "🟥", "🟥", "🟩", "🟨"}, false),
	Entry("AABBB BBAAA", "AABBB", "BBAAA", [5]string{"🟨", "🟨", "🟨", "🟨", "🟥"}, false),
	Entry("AABBB BBAAB", "AABBB", "BBAAB", [5]string{"🟨", "🟨", "🟨", "🟨", "🟩"}, false),
	Entry("AABBB BBACA", "AABBB", "BBACA", [5]string{"🟨", "🟨", "🟨", "🟥", "🟨"}, false),
	Entry("ABABA BABAB", "ABABA", "BABAB", [5]string{"🟨", "🟨", "🟨", "🟨", "🟥"}, false),
)

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
		{"HELLO HELLO", args{"HELLO", "HELLO"}, [5]string{"🟩", "🟩", "🟩", "🟩", "🟩"}, true},
		{"PARTS DARTS", args{"PARTS", "DARTS"}, [5]string{"🟥", "🟩", "🟩", "🟩", "🟩"}, false},
		{"HACKS SACKS", args{"HACKS", "SACKS"}, [5]string{"🟥", "🟩", "🟩", "🟩", "🟩"}, false},
		{"EAGER EAGLE", args{"EAGER", "EAGLE"}, [5]string{"🟩", "🟩", "🟩", "🟥", "🟨"}, false},
		{"EARNS EAGLE", args{"EARNS", "EAGLE"}, [5]string{"🟩", "🟩", "🟥", "🟥", "🟥"}, false},
		{"STILL LOLLS", args{"STILL", "LOLLS"}, [5]string{"🟨", "🟥", "🟥", "🟩", "🟨"}, false},
		{"AABBB BBAAA", args{"AABBB", "BBAAA"}, [5]string{"🟨", "🟨", "🟨", "🟨", "🟥"}, false},
		{"AABBB BBAAB", args{"AABBB", "BBAAB"}, [5]string{"🟨", "🟨", "🟨", "🟨", "🟩"}, false},
		{"AABBB BBACA", args{"AABBB", "BBACA"}, [5]string{"🟨", "🟨", "🟨", "🟥", "🟨"}, false},
		{"ABABA BABAB", args{"ABABA", "BABAB"}, [5]string{"🟨", "🟨", "🟨", "🟨", "🟥"}, false},
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
