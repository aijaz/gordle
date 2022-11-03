package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWordleTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WordleTest Suite")
}
