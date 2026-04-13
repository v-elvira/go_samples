package main

import "fmt"

type counter struct {
	value int
}

func (c *counter) increment() { c.value++ }

type usage struct {
	service string
	counter
}

func main() {

	u := new(usage)
	u.increment()
	u.increment()
	fmt.Println(u)       // &{ {2}}
	fmt.Println(u.value) // 2

}
