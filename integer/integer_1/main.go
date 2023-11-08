package main

import "fmt"

func main() {
	var distance int // в сантиметрах

	fmt.Print("Введите расстояние:")

	fmt.Scan(&distance)

	convertedDistance := distance / 100

	fmt.Printf("Кол-во полных метров: %d", convertedDistance)
}
