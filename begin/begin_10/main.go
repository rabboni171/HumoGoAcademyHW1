package main

import "fmt"

func main() {
	var numberOne, numberTwo int

	fmt.Println("Введите два числа:")

	fmt.Scan(&numberOne, &numberTwo)

	sumOfSquares := numberOne*numberOne + numberTwo*numberTwo

	diffOfSquares := numberOne*numberOne - numberTwo*numberTwo

	multipOfSquares := numberOne * numberOne * numberTwo * numberTwo

	divisionOfSquares := numberOne * numberOne / numberTwo * numberTwo

	fmt.Printf("Сумма квадратов: %d\n", sumOfSquares)

	fmt.Printf("Разность квадратов: %d\n", diffOfSquares)

	fmt.Printf("Произведение квадратов: %d\n", multipOfSquares)

	fmt.Printf("Частное квадратов: %d", divisionOfSquares)

}
