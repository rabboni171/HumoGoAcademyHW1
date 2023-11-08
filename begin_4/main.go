package main

import (
	"fmt"
	"math"
)

func main() {
	var diameterOfCircle float64

	fmt.Println("Введите диаметр окружности:")

	fmt.Scan(&diameterOfCircle)

	fmt.Printf("Длина окружности: %f", math.Pi*float64(diameterOfCircle))
}
