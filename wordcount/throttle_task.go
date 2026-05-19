// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	ch_cancel := make(chan struct{})

	ticker := time.NewTicker(time.Duration(1000/limit) * time.Millisecond)

	cancel = func() {
		select {
		case <-ch_cancel:
		default:
			ticker.Stop()
			close(ch_cancel)
		}
	}

	handle = func() error {
		select {
		case <-ch_cancel:
			return ErrCanceled
		case <-ticker.C:
			go fn()
			return nil
		}
	}
	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	cancel()
	handle()
	handle()
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
