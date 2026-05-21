package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Потеря точности ("испарение" единицы)
	var x float64 = 1e25
	fmt.Println("x =", x)
	fmt.Println("x + 1 =", x+1)                 // Выведет то же самое число! 1 исчезла.
	fmt.Println("Равны ли x и x+1?:", x == x+1) // Выведет: true

	fmt.Println("\n--- Эксперимент с ростом мантиссы ---")

	// 2. Растет ли мантисса сама по себе при умножении?
	var num float64 = 1.23456789012345 // Красивая мантисса
	for i := 0; i < 5; i++ {
		// Получаем биты числа, чтобы заглянуть "внутрь" процессора
		bits := math.Float64bits(num)
		// Маска для извлечения 52 бит мантиссы
		mantissaBits := bits & 0x000FFFFFFFFFFFFF

		fmt.Printf("Число: %e | Мантисса (в hex): %x\n", num, mantissaBits)
		num *= 10
	}
	fmt.Println(100 < math.Inf(-1), 100.5 < math.Inf(1), math.Inf(1))
	fmt.Println(100 < 101.5)
	// a, b := 100, 100.5
	// fmt.Println(a < b)  // invalid operation: a < b (mismatched types int and float64)
	fmt.Printf("%T\n", math.Inf(1))
	// c, d := 100, math.Inf(1)
	// fmt.Println(c < d)  // (mismatched types int and float64)
	v := math.Inf(-1)
	fmt.Println(math.IsInf(v, -1), math.IsInf(v, 0), math.IsInf(v, 42)) //-:-Inf  0:Inf  +:+Inf
}
