package main

import "fmt"

func main() {
	var numberOne, numberTwo float64

	fmt.Println("Введите числа:")

	fmt.Scan(&numberOne, &numberTwo)

	avarageArithmetic := (numberOne + numberTwo) / 2

	fmt.Printf("Cреднее арифметическое: %.2f", avarageArithmetic)
}
