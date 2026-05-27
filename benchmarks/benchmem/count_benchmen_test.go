package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Regexp-%d", length)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountRegexp(phrase)
			}
		})
	}
}

func BenchmarkFields(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Fields-%d", length)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountFields(phrase)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountSplit(phrase)
			}
		})
	}
}

func BenchmarkLowerPhr(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("LowerPhr-%d", length)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountLowerPhrase(phrase)
			}
		})
	}
}

func BenchmarkAllocate(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				WordCountAllocate(phrase)
			}
		})
	}
}

// randomPhrase возвращает фразу из n случайных слов.
func randomPhrase(n int) string {
	words := make([]string, n)
	for i := range words {
		words[i] = randomWord(3)
	}
	return strings.Join(words, " ")
}

var rnd = rand.New(rand.NewSource(0))

// randomWord возвращает слово из n случайных букв.
func randomWord(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rnd.Intn(len(letters))]
	}
	return string(chars)
}
