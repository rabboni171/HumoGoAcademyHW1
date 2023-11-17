package main

import "fmt"

func main() {
	var price float64
	fmt.Scan(&price)
	for i := 1; i <= 10; i++ {
		anotherMass := float64(i) / 10.0
		totalPrice := price * anotherMass
		fmt.Printf("%.1f кг: %.1f\n", anotherMass, totalPrice)
	}

}
