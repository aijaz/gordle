package wordle

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidFiveLetterWordIsValid(t *testing.T) {
	valid := isGuessLengthValid("ocean")
	assert.True(t, valid, "The word is valid")
	assert.False(t, isGuessLengthValid("Foo"), "The word is not valid")
}

func TestValidThreeLetterWordIsNotValid(t *testing.T) {
	valid := isGuessLengthValid("foo")
	assert.False(t, valid, "The word is not valid")
}

func TestWordleCheckTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wordle Suite")
}

var _ = Describe("Wordle", func() {
	It("Should work", func() {
		Expect("foo").To(Equal("foo"))
	})
})
