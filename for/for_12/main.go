package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	multiple := 1.0
	for i := 1.1; i <= float64(n); i++ {
		multiple *= i

	}
	fmt.Println(multiple)
}
