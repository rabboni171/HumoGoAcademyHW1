package main

import (
	"fmt"
	"math"
)

func main() {
	var numberOne, numberTwo float64

	fmt.Println("Введите числа:")

	fmt.Scan(&numberOne, &numberTwo)

	avarageGeometric := math.Sqrt(numberOne * numberTwo)
	
	fmt.Printf("Cреднее геометрическое: %.1f", avarageGeometric)
}
