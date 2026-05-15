package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("generate done")
		defer close(out)
		for {
			word := randomWord(5)
			select {
			case out <- word:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func isUnique(val string) bool {
	seen := make(map[rune]bool)
	for _, r := range val {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

type pair struct {
	word string
	rev  string
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer fmt.Println("takeUnique done")
		defer close(out)
		for {
			select {
			case val := <-in:
				if isUnique(val) {
					select {
					case out <- val:
					case <-cancel:
						return
					}
				}
			case <-cancel:
				return

			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan pair {
	out := make(chan pair)
	go func() {
		defer fmt.Println("reverse done")
		defer close(out)
		for {
			select {
			case val := <-in:
				runes := []rune(val)
				n := len(val)
				for i := 0; i < n/2; i++ {
					runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
				}
				select {
				case out <- pair{val, string(runes)}:
				case <-cancel:
					return
				}

			case <-cancel:
				return
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan pair) <-chan pair {
	out := make(chan pair)
	go func() {
		defer fmt.Println("merge done")
		defer close(out)
		for c1 != nil || c2 != nil {
			select {
			case v1, ok := <-c1:
				if ok {
					select {
					case out <- v1:
					case <-cancel:
						return
					}
				} else {
					c1 = nil
				}
			case v2, ok := <-c2:
				if ok {
					select {
					case out <- v2:
					case <-cancel:
						return
					}
				} else {
					c2 = nil
				}
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan pair, n int) {
	defer fmt.Println("print done")
	for range n {
		pair := <-in
		fmt.Printf("%s -> %s\n", pair.word, pair.rev)
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	//defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
	close(cancel)
	time.Sleep(50 * time.Millisecond)
}

// ERROR: there are leaked goroutines (fixed with nested select)
