package main

import "fmt"

func main() {
	for i := range 3 {
		fmt.Println(i)
	}
	// fmt.Println(i) // undefined i
	var r rune = '🚀'
	fmt.Printf("Rocket: %d\n", r)
	var t int8 = 127
	t += 100
	fmt.Println(t) // -29

	var value int64 = 300

	// Опасно
	fmt.Println(int8(value))         // 44 потеря данных
	fmt.Println(52 | 100)            // 116
	fmt.Printf("Error: %05d\n", 404) // Error: 00404
	result := fmt.Sprintf("Hex: %X, Bin: %b", 255, 255)
	fmt.Println(result)
	fmt.Print("A", "В", "С")
	fmt.Println("D", "Е", "F")
	fmt.Printf("res: %+08.2f\n", 3.148) // res: +0003.15
	fmt.Printf("|%5d|\n", 123)          // | 123 |
	fmt.Printf("|%-5d|\n", 456)         // |456  |
}
