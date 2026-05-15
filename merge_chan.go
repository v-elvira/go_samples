package main

import (
	"fmt"
	"sync"
	"time"
)

// rangeGen отправляет в канал числа от start до stop-1
func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}

	}()
	return out
}

// начало решения

// merge выбирает числа из входных каналов и отправляет в выходной
func merge(channels ...<-chan int) <-chan int {
	// объедините все исходные каналы в один выходной
	// последовательное объединение НЕ подходит
	merged := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Go(func() {
			for item := range ch {
				merged <- item
			}
		})

	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}

// конец решения

func main() {
	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)
	in3 := rangeGen(31, 35)

	start := time.Now()
	merged := merge(in1, in2, in3)
	for val := range merged {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("Took", time.Since(start))
}
