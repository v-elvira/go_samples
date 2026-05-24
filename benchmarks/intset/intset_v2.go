package main

// IntSet2 реализует множество целых чисел
// (элементы множества уникальны).
type IntSet2 struct {
	vals map[int]struct{}
}

// MakeIntSet создает пустое множество.
func MakeIntSet2() IntSet2 {
	return IntSet2{vals: make(map[int]struct{})}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet2) Contains(elem int) bool {
	_, exists := s.vals[elem]
	return exists
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet2) Add(elem int) bool {
	_, exists := s.vals[elem]
	if exists {
		return false
	}
	s.vals[elem] = struct{}{}
	return true
}
