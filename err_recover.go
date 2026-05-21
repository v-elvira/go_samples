package main

import "fmt"

// getChar возвращает символ строки по указанному индексу
func getChar(str string, idx int) byte {
	return str[idx]
}

func main() {
	defer func() {
		recover() // =one(!) level inline to work
	}()
	c := getChar("hello", 10)
	//panic: runtime error: index out of range [10] with length 5
	fmt.Println("hello[10] = ", c)

	// ручная паника
	panic("oops")
}
