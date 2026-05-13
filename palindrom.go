package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome_v1(str string) bool {
	lower := []rune(strings.ToLower(str))
	n := len(lower)
	for i := range n / 2 {
		if lower[i] != lower[n-1-i] {
			return false
		}
	}
	return true
}

func IsPalindrome(str string) bool {
	lower := []rune(strings.ToLower(str))

	for l, r := 0, len(lower)-1; l < r; {
		if unicode.IsSpace(lower[l]) || unicode.IsPunct(lower[l]) {
			l++
			continue
		}
		if unicode.IsSpace(lower[r]) || unicode.IsPunct(lower[r]) {
			r--
			continue
		}
		if lower[l] != lower[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(("Aha")))
	fmt.Println(IsPalindrome(("Приирп")))
	fmt.Println(IsPalindrome(("ф ф")))
}
