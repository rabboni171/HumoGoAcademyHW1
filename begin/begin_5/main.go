package main

import (
	"fmt"
	"math"
)

func main() {
	var cubeEdge int

	fmt.Println("Введите ребро куба:")
	
	fmt.Printf("Объем куба: %.0f", math.Pow(float64(cubeEdge), 3))
	
	fmt.Scan(&cubeEdge)

	fmt.Printf("Объем куба: %.0f\n", math.Pow(float64(cubeEdge), 3))

	fmt.Printf("Площадь поверхности куба: %.0f", 6*(math.Pow(float64(cubeEdge), 2)))
}
