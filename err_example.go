package main

import (
	"fmt"
	"strings"
)

// // фрагмент кода стандартной библиотеки
// type error interface {
//     Error() string
// }

type lookupError struct {
	src    string
	substr string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)
}

func indexOf(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		// Создаем и возвращаем ошибку типа `lookupError`.
		return -1, lookupError{src, substr}
	}
	return idx, nil
}

func main() {
	src := "go is awesome"
	for _, substr := range []string{"go", "js"} {
		if res, err := indexOf(src, substr); err != nil {
			fmt.Printf("indexOf(%#v, %#v) failed: %v\n", src, substr, err)
		} else {
			fmt.Printf("indexOf(%#v, %#v) = %v\n", src, substr, res)
		}
	}
	// indexOf("go is awesome", "go") = 0
	// indexOf("go is awesome", "js") failed: 'js' not found in 'go is awesome'

	fmt.Println("--------------------------")

	_, err := indexOf(src, "js")
	if err, ok := err.(lookupError); ok {
		fmt.Println("err.src:", err.src)
		fmt.Println("err.substr:", err.substr)
	}
	// err.src: go is awesome
	// err.substr: js

}
