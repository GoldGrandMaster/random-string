package randomstring

import (
	"math/rand"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-")

type rs struct {
	letters []rune
}

func (r *rs) randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = r.letters[rand.Intn(len(r.letters))]
	}
	return string(b)
}

// New generates a new randomstring struct
func New() *rs {
	return &rs{
		letters: letters,
	}
}

// Letters set the letters of random string
func (r *rs) Letters(letters string) *rs {
	r.letters = []rune(letters)
	return r
}

// Generate a random string with the given length
func (r *rs) Generate(length int) string {
	rand.Seed(time.Now().UnixNano())
	return r.randSeq(length)
}
