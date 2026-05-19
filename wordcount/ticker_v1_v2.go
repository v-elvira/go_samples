package main

import (
	"fmt"
	"time"
)

// начало решения v1 (my)
func schedule(dur time.Duration, fn func()) func() {
	ch_cancel := make(chan struct{})
	ch_token := make(chan int, 1)
	ch_token <- 1

	cancel := func() {
		select {
		case <-ch_cancel:
			return
		default:
			ch_cancel <- struct{}{}
		}
	}

	go func() {
		defer close(ch_token)
		for {
			time.Sleep(dur)
			select {
			case <-ch_cancel:
				close(ch_cancel)
				return
			default:
				select {
				case <-ch_token:
					go func() {
						fn()
						ch_token <- 1
					}()
				default:
				}
			}
		}
	}()

	return cancel
}

// конец решения

// // v2 (NewTicker) start

// func schedule(dur time.Duration, fn func()) func() {
// 	ticker := time.NewTicker(dur)
// 	canceled := make(chan struct{})

// 	tick := func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				fn()
// 			case <-canceled:
// 				return
// 			}
// 		}
// 	}

// 	cancel := func() {
// 		select {
// 		case <-canceled:
// 			return
// 		default:
// 			ticker.Stop()
// 			close(canceled)
// 		}
// 	}

// 	go tick()
// 	return cancel
// }

// // end v2

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
