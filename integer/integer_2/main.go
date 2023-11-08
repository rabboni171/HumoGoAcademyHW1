package main

import "fmt"

func main() {
	var mass int // в кг

	fmt.Println("Введите массу:")

	fmt.Scan(&mass)

	convertedMass := mass / 1000

	fmt.Printf("Кол-во полных тонн: %d", convertedMass)
}
