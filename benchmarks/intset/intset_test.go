package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkIntSet(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		set := randomSet(size)
		name := fmt.Sprintf("Contains-%d", size)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				elem := rand.Intn(100000)
				set.Contains(elem)
			}
		})

	}
}

var rnd = rand.New(rand.NewSource(42))

func randomSet(size int) IntSet {
	set := MakeIntSet()
	for i := 0; i < size; i++ {
		n := rnd.Intn(100000)
		set.Add(n)
	}
	return set
}

// ----------------------------------------
// v2:
func BenchmarkIntSet2(b *testing.B) {
	fmt.Println("v2------------------------------------------")
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		set := randomSet2(size)
		name := fmt.Sprintf("Contains-%d", size)
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				elem := rand.Intn(100000)
				set.Contains(elem)
			}
		})

	}
}

func randomSet2(size int) IntSet2 {
	set := MakeIntSet2()
	for i := 0; i < size; i++ {
		n := rnd.Intn(100000)
		set.Add(n)
	}
	return set
}
