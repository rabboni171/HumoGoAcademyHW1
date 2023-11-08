package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Введите числа A и B:")

	fmt.Scan(&a, &b)

	chisloOtrezkov := a / b

	fmt.Printf("Кол-во отрезко B: %d", chisloOtrezkov)

}
