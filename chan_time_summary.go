package main

import (
	"errors"
	"fmt"
	"time"
)

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	var result int

	done := make(chan struct{})
	go func() {
		defer close(done)
		result = fn()
	}()

	select {
	case <-done:
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

//--------------------------------------------------
timer := time.NewTimer(100 * time.Millisecond)
go func() {
    eventTime = <-timer.C
    work()
}()

if !timer.Stop() {
    fmt.Println("too late to cancel")
}

//--------------------------------------------------
work := func() {
    fmt.Println("work done")
}

time.AfterFunc(100*time.Millisecond, work)



//--------------------------------------------------
ticker := time.NewTicker(50 * time.Millisecond)
defer ticker.Stop()

go func() {
    for {
        at := <-ticker.C
        work(at)
    }
}()
