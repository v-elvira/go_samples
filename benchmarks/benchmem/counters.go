package main

import (
	"regexp"
	"strings"
)

// Counter хранит количество вхождений слов.
type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

// WordCountRegexp считает абсолютные частоты слов в строке.
// Использует Regexp.Split() для разделения строки на слова.
func WordCountRegexp(s string) Counter {
	counter := make(Counter)
	for _, word := range splitter.Split(s, -1) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

// WordCountFields считает абсолютные частоты слов в строке.
// Использует strings.Fields() для разделения строки на слова.
func WordCountFields(s string) Counter {
	counter := make(Counter)
	for _, word := range strings.Fields(s) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

// WordCountSplit считает абсолютные частоты слов в строке.
// Использует strings.Split() для разделения строки на слова.
func WordCountSplit(s string) Counter {
	counter := make(Counter)
	for _, word := range strings.Split(s, " ") {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

// WordCountLowerPhrase считает абсолютные частоты слов в строке.
// Преобразует всю строку к нижнему регистру перед разделением.
func WordCountLowerPhrase(s string) Counter {
	counter := make(Counter)
	phrase := strings.ToLower(s)
	for _, word := range strings.Split(phrase, " ") {
		counter[word]++
	}
	return counter
}

// WordCountAllocate считает абсолютные частоты слов в строке.
// Предварительно выделяет память для счетчика.
func WordCountAllocate(s string) Counter {
	words := strings.Split(s, " ")
	size := len(words) / 2
	if size > 10000 {
		size = 10000
	}
	counter := make(Counter, size)
	for _, word := range words {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}
