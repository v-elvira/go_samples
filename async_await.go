package main

import (
	"fmt"
	"time"
)

// async превращает переданную функцию в асинхронную.
// Асинхронная функция при вызове возвращает канал,
// в котором появится результат.
func async(fn func() any) func() <-chan any {
	return func() <-chan any {
		done := make(chan any, 1)
		go func() {
			done <- fn()
		}()
		return done
	}
}

// await ожидает результата выполнения асинхронной функции
// на переданном канале.
func await(in <-chan any) any {
	return <-in
}

func main() {
	fn := func() any {
		time.Sleep(500 * time.Millisecond)
		// do stuff
		return "okay"
	}

	slowpoke := async(fn) // делаем асинхронную функцию

	start := time.Now()
	slowpoke()                  // не блокирует
	slowpoke()                  // не блокирует
	slowpoke()                  // не блокирует
	result := await(slowpoke()) // блокирует до получения результата

	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println("took", elapsed)
	// okay
	// took 500ms

	// общее время выполнения 500 мс, а не 2 с!
}
