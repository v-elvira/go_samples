package main

import (
	"strings"
	// "testing"
	"fmt"
	"unicode"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	filter := func(c rune) rune {
		switch {
		case unicode.IsLetter(c):
			return unicode.ToLower(c)
		case unicode.IsDigit(c):
			return c
		}
		return '-'
	}

	return strings.Map(filter, src)
}

// конец решения

func main() {
	const phrase = "Go - Is - Awesome"
	const want = "go---is---awesome"
	got := slugify(phrase)
	fmt.Printf("%s: got %#v, want %#v", phrase, got, want)
}

// func Test(t *testing.T) {
// 	const phrase = "Go Is Awesome!"
// 	const want = "go-is-awesome"
// 	got := slugify(phrase)
// 	if got != want {
// 		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
// 	}
// }
