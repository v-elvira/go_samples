// -- more/more.go --
package main

import (
	"fmt"

	"github.com/v-elvira/go_samples/more/numbers"
	"github.com/v-elvira/go_samples/more/text"
)

func main() {
	ok := text.ArePermutation("hello", "lehol")
	fmt.Println("hello / lehol →", ok)

	ok = numbers.IsEven(42)
	fmt.Println("42 is even →", ok)
}
