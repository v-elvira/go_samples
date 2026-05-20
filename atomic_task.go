package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// начало решения
// V1: all atomic.Int32 methods visible

// type Total struct {
// 	atomic.Int32
// } // anon field => embedded (all atomic.Int64 methods inherited)

// func (t *Total) Increment() {
// 	t.Add(1)
// }

// func (t *Total) Value() int {
// 	return int(t.Load())
// }
// // конец решения

// начало решения V2
type Total struct {
	value atomic.Int32
} // anon field => embedded (all atomic.Int64 methods inherited)

func (t *Total) Increment() {
	t.value.Add(1)
}

func (t *Total) Value() int {
	return int(t.value.Load())
}

// конец решения

func main() {
	var wg sync.WaitGroup

	var total Total

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				total.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("total", total.Value())
}
