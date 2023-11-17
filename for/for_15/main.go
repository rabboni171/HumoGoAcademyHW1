package main

import (
	"fmt"
)

func main() {
	var a float64
	var n int
	result := 1.0
	fmt.Scan(&a, &n)

	for i := 0; i < n; i++ {
		result *= a
		// result = result * a
	}

	fmt.Println(result)
}
