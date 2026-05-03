// -- more/numbers/numbers.go --
package numbers

import "github.com/v-elvira/go_samples/more_err/text" // import cycles not allowed

// IsEven проверяет число на четность.
// IsEven проверяет число на четность.
func IsEven(n int) bool {
	return n%2 == 0
}

// AsDigits возвращает срез с цифрами числа.
func AsDigits(n int) []int {
	runes := text.AsRunes(n) // text.AsRunes(n)
	count := len(runes)
	zero := int('0')
	digits := make([]int, count)
	for idx, char := range runes {
		digits[idx] = int(char) - zero
	}
	return digits
}
