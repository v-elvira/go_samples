package main

import (
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	result := 0
	for _, phrase := range directions {
		words := strings.Fields(phrase)
		for _, word := range words {
			if !unicode.IsDigit(rune(word[0])) {
				continue
			}
			// fmt.Println(word)
			if strings.HasSuffix(word, "km") {
				val, _ := strconv.ParseFloat(word[:len(word)-2], 64)
				result += int(val * 1000)
			} else if strings.HasSuffix(word, "m") {
				val, _ := strconv.Atoi(word[:len(word)-1])
				result += val
			}
		}
	}
	return result
}

// конец решения

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	const want = 6000
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}
