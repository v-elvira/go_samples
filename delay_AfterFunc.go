package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	ch_cancel := make(chan struct{})

	cancel := func() {
		select {
		case <-ch_cancel: // closed 0-val
		default:
			ch_cancel <- struct{}{}
		}
	}

	go func() {
		select {
		case <-ch_cancel:
			close(ch_cancel)
			return
		case <-time.After(dur):
			close(ch_cancel)
			fn()
		}
	}()
	return cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
