package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 { // os.Args[0] — это путь к самой программе
		fmt.Println(os.Args[0])
		fmt.Println("Ошибка: передайте строку в кавычках.")
		return
	}
	fmt.Println(len(strings.Fields(strings.TrimSpace(os.Args[1]))))
}
