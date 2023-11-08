package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Println("Введите ребрa параллелепипеда:")

	fmt.Scan(&a, &b, &c)

	fmt.Printf("Объем прда: %d\n", a*b*c)

	fmt.Printf("Площадь поверхности прда: %d", 2*(a*b+b*c+a*c))
}
