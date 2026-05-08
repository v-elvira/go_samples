// Сколько цифр в каждом слове?
package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// counter хранит количество цифр в каждом слове.
// Ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// начало решения

// countDigitsInWords считает количество цифр в словах фразы.
func countDigitsInWords(phrase string) counter {
	var wg sync.WaitGroup
	syncStats := new(sync.Map)
	words := strings.Fields(phrase)

	// Посчитайте количество цифр в словах,
	// используя отдельную горутину для каждого слова.

	// Чтобы записать результаты подсчета,
	// используйте syncStats.Store(word, count)

	for _, word := range words {
		wg.Go(func() {
			syncStats.Store(word, countDigits(word))
		})
	}
	wg.Wait()

	// В результате syncStats должна содержать слова
	// и количество цифр в каждом.

	return asStats(syncStats)
}

// конец решения

// countDigits возвращает количество цифр в строке.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// asStats преобразует статистику из sync.Map в обычную карту.
func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStats печатает количество цифр в словах.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)
}
