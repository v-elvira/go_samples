package main

import "fmt"

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
type IntSet struct {
	elems *[]int
}

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	elems := []int{}
	return IntSet{&elems}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
	for _, el := range *s.elems {
		if el == elem {
			return true
		}
	}
	return false
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	*s.elems = append(*s.elems, elem)
	return true
}

func main() {
	set := MakeIntSet()

	set.Add(5)
	fmt.Println(set.Contains(5))
	// true

	fmt.Println(set.Contains(42))
	// false

	// элементы множества уникальны,
	// добавить дважды один и тот же элемент не получится
	added := set.Add(5)
	fmt.Println(added)
	// false
}
