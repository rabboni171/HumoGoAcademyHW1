package main

import "fmt"

func main() {
	var price float64
	fmt.Scan(&price)
	for i := 1; i <= 10; i++ {
		result := price * float64(i)
		fmt.Printf("%d кг: %.1f\n", i, result)
	}
}
