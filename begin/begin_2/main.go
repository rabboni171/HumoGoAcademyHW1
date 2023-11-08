package main

import "fmt"

func main() {
	var sideOfSquare int

	fmt.Println("Введите сторону квадрата:")

	fmt.Scan(&sideOfSquare)

	fmt.Printf("Площадь квадрата: %d", sideOfSquare*sideOfSquare)
}
