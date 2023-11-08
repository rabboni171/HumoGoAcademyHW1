package main

import "fmt"

func main() {
	var lengthOfRectangle int

	var widthOfRectangle int

	fmt.Println("Введите стороны прямоугольника")

	fmt.Scan(&lengthOfRectangle, &widthOfRectangle)

	fmt.Printf("Площадь: %d\n", lengthOfRectangle*widthOfRectangle)

	fmt.Printf("Периметр: %d", 2*(lengthOfRectangle+widthOfRectangle))

}
