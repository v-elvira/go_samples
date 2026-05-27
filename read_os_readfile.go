package main

import (
	"fmt"
	"os"
	"strings"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	strs := strings.Split(string(data), "\n")
	if len(strs) > 0 && strs[len(strs)-1] == "" {
		strs = strs[:len(strs)-1]
	}
	return strs, nil
}

// конец решения

func main() {
	lines, err := readLines("/etc/passwd")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
