// Promise.all()
package main

import (
	"fmt"
	"time"
)

// начало решения

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// Выполните все переданные функции,
	// соберите результаты в срез и верните его.
	type pair struct {
		ind    int
		result any
	}
	ch := make(chan pair)
	for i, f := range funcs {
		go func() {
			ch <- pair{i, f()}
		}()
	}
	result := make([]any, len(funcs))
	for range len(funcs) {
		p := <-ch
		result[p.ind] = p.result
	}
	return result
}

// конец решения

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(1), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
