package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// начало решения

// ErrFailed и ErrManual - причины остановки цикла.
var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

// Worker выполняет заданную функцию в цикле, пока не будет остановлен.
// Гарантируется, что Worker используется только в одной горутине.
type Worker struct {
	fn func() error
	// TODO: добавить поля
	ctx     context.Context
	cancel  func(error)
	started bool
}

// NewWorker создает новый экземпляр Worker с заданной функцией.
// Но пока не запускает цикл с функцией.
func NewWorker(fn func() error) *Worker {
	ctx, cancel := context.WithCancelCause(context.Background())
	return &Worker{fn: fn, ctx: ctx, cancel: cancel}
}

// Start запускает отдельную горутину, в которой циклически
// выполняет заданную функцию, пока не будет вызван метод Stop,
// либо пока функция не вернет ошибку.
// Повторные вызовы Start игнорируются.
func (w *Worker) Start() {
	if w.started {
		return
	}
	w.started = true
	// TODO: реализовать требования
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
				err := w.fn()
				if err != nil {
					w.cancel(ErrFailed)
					return
				}
			}
		}
	}()
}

// Stop останавливает выполнение цикла.
// Вызов Stop до Start игнорируется.
// Повторные вызовы Stop игнорируются.
func (w *Worker) Stop() {
	// TODO: реализовать требования
	w.cancel(ErrManual)
}

// AfterStop регистрирует функцию, которая
// будет вызвана после остановки цикла.
// Можно зарегистрировать несколько функций.
// Вызовы AfterStop после Start игнорируются.
func (w *Worker) AfterStop(fn func()) {
	// TODO: реализовать требования
	if w.started {
		return
	}
	context.AfterFunc(w.ctx, fn)
}

// Err возвращает причину остановки цикла:
// - ErrManual - вручную через метод Stop;
// - ErrFailed - из-за ошибки, которую вернула функция.
func (w *Worker) Err() error {
	// TODO: реализовать требования
	return context.Cause(w.ctx)
}

// конец решения

func main() {
	{
		// Start-Stop
		count := 9
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(105 * time.Millisecond)
		worker.Stop()

		fmt.Println()
		// 9 8 7 6 5 4 3 2 1 0
	}
	{
		// ErrFailed
		count := 3
		fn := func() error {
			fmt.Print(count, " ")
			count--
			if count == 0 {
				return errors.New("count is zero")
			}
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		worker.Start()
		time.Sleep(35 * time.Millisecond)
		worker.Stop()

		fmt.Println(worker.Err())
		// 3 2 1 failed
	}
	{
		// AfterStop
		fn := func() error { return nil }

		worker := NewWorker(fn)
		worker.AfterStop(func() {
			fmt.Println("called after stop")
		})

		worker.Start()
		worker.Stop()

		time.Sleep(10 * time.Millisecond)
		// called after stop
	}
}

// // V2 (from solution):

// // ErrFailed и ErrManual - причины остановки цикла.
// var ErrFailed = errors.New("failed")
// var ErrManual = errors.New("manual")

// // Worker выполняет заданную функцию в цикле, пока не будет остановлен.
// type Worker struct {
//     fn         func() error
//     ctx        context.Context
//     cancel     func(cause error)
//     afterFuncs []func()
// }

// // NewWorker создает новый экземпляр Worker с заданной функцией.
// func NewWorker(fn func() error) *Worker {
//     return &Worker{fn: fn}
// }

// // Start запускает отдельную горутину, в которой циклически
// // выполняет заданную функцию, пока не будет вызван метод Stop,
// // либо пока функция не вернет ошибку.
// func (w *Worker) Start() {
//     if w.ctx != nil {
//         return
//     }
//     w.ctx, w.cancel = context.WithCancelCause(context.Background())
//     for _, fn := range w.afterFuncs {
//         context.AfterFunc(w.ctx, fn)
//     }
//     go w.work()
// }

// // Stop останавливает выполнение цикла.
// func (w *Worker) Stop() {
//     if w.ctx == nil {
//         return
//     }
//     w.cancel(ErrManual)
// }

// // AfterStop регистрирует функцию, которая
// // будет вызвана после остановки цикла.
// func (w *Worker) AfterStop(fn func()) {
//     if w.ctx != nil {
//         return
//     }
//     w.afterFuncs = append(w.afterFuncs, fn)
// }

// // Err возвращает причину остановки цикла.
// func (w *Worker) Err() error {
//     return context.Cause(w.ctx)
// }

// // work выполняет заданную функцию в цикле.
// func (w *Worker) work() {
//     for {
//         select {
//         case <-w.ctx.Done():
//             return
//         default:
//             err := w.fn()
//             if err != nil {
//                 w.cancel(ErrFailed)
//                 return
//             }
//         }
//     }
// }
