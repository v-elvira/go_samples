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
	filter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		// return !(c >= 'a' && c <= 'z') && !unicode.IsNumber(c) && !(c == '-') // PASSED
	}
	words := strings.FieldsFunc(strings.ToLower(src), filter)
	return strings.Join(words, "-")
}

// конец решения

func main() {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
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
