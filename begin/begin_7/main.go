package main

import (
	"fmt"
	"math"
)

func main() {
	var RadiusOfCircle float64

	fmt.Println("Введите радиус окружности:")

	fmt.Scan(&RadiusOfCircle)

	fmt.Printf("Длина окружности: %.1f\n", 2*math.Pi*RadiusOfCircle)

	fmt.Printf("Площадь круга: %.1f", math.Pi*RadiusOfCircle*RadiusOfCircle)

}
