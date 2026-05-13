package main

import (
	"fmt"
	"strings"
)

func numberWithSeparators(s string) string {
	var b strings.Builder
	b.Grow(len(s) + len(s)/3)
	n := len(s) % 3
	if n == 0 {
		n = 3
	} else if n == 1 && s[0] == '-' {
		n = 4
	}
	b.WriteString(s[:n])
	for ; n < len(s); n += 3 {
		b.WriteByte('.')
		b.WriteString(s[n : n+3])
	}
	return b.String()
}

func main() {
	fmt.Println(numberWithSeparators("2412145"))
	fmt.Println(numberWithSeparators("214"))
	fmt.Println(4 / 3)
	fmt.Println(numberWithSeparators("123123"))
	fmt.Println(numberWithSeparators("-123123"))
	fmt.Println(numberWithSeparators("-12312"))
	fmt.Println(numberWithSeparators("-1231"))
	fmt.Println(numberWithSeparators("-1231234"))
	fmt.Println(numberWithSeparators("-123"))
}
