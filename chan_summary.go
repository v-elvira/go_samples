package main

import "fmt"

// cancel channel:
func send(cancel chan struct{}, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)
	ch := send(cancel, 100)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//--------------------------------------------------

// // result channel:
// func send(n int) <-chan int {
//     out := make(chan int)
//     go func() {
//         defer close(out)
//         for i := 1; i <= n; i++ {
//             out <- i
//         }
//     }()
//     return out
// }

// func main() {
//     ch := send(5)
//     for n := range ch {
//         fmt.Print(n)
//     }
// }

//--------------------------------------------------

// // done channel:
// func work() <-chan struct{} {
//     done := make(chan struct{})
//     go func() {
//         defer close(done)
//         fmt.Println("work done")
//     }()
//     return done
// }

// func main() {
//     done := work()
//     <-done
// }
