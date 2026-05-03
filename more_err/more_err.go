// -- more/more.go --
package main

import (
	"fmt"
	"github.com/v-elvira/go_samples/more_err/numbers"
	// "github.com/v-elvira/go_samples/more_err/text"
)

func main() {
	digits := numbers.AsDigits(42513)
	fmt.Println("42513 →", digits)
}
