package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Reader:
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя с пробелами: ")
	fullName, err := reader.ReadString('\n') // read till \n {include first \n}
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fullName = strings.TrimSpace(fullName)

	fmt.Printf("Результат Reader: >>>%s<<<\n", fullName)

	// Scanner:
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите имя: ")
	scanner.Scan()
	fullName = scanner.Text()
	fullName = strings.TrimSpace(fullName) // optional, for " " {no \n in the end}

	fmt.Printf("Результат Scanner: >>>%s<<<\n", fullName)
}
