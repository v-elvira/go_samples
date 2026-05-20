package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	vals map[string]int
	lock sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	c.vals[str]++
	c.lock.Unlock()
}

func (c *Counter) Value(str string) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.vals[str]
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.Lock()
	defer c.lock.Unlock()
	for k, v := range c.vals {
		fn(k, v)
	}

}

func NewCounter() *Counter {
	c := Counter{vals: make(map[string]int), lock: sync.Mutex{}}
	return &c
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
