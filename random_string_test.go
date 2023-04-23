package randomstring_test

import (
	"github.com/theyahya/random-string"
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	length := 10
	randomString := randomstring.New().Generate(10)
	expected := len(randomString)
	if length != expected {
		t.Errorf("expected %v", expected)
	}
}

func TestRandomStringLetters(t *testing.T) {
	randomString := randomstring.New().Letters("a").Generate(2)
	expected := "aa"
	if randomString != expected {
		t.Errorf("expected %v", expected)
	}
}
