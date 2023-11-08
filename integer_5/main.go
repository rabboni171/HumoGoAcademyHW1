package main

import "fmt"

func main() {
	var a, b, remainder int

	fmt.Print("Введите значение a: ")

	fmt.Scanln(&a)

	fmt.Print("Введите значение b: ")

	fmt.Scanln(&b)

	remainder = a % b

	fmt.Println("Длина незанятой части отрезка a:", remainder)
}
