package main

import (
	"fmt"
	"strings"
)

// начало решения

func slugify(src string) string {
	var b strings.Builder
	b.Grow(len(src))
	space := false
	for _, item := range src {
		if item >= 65 && item <= 90 {
			if space {
				b.WriteByte(45)
				space = false
			}
			b.WriteByte(byte(item + 32))
		} else if item == 45 || (item >= 48 && item <= 57) || (item >= 97 && item <= 122) {
			if space {
				b.WriteByte(45)
				space = false
			}
			b.WriteByte(byte(item))
		} else if b.Len() > 0 {
			space = true
		}
	}
	return b.String()
}

// конец решения

func main() {
	const phrase = "A 100x Investment (2019)"
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019
}
