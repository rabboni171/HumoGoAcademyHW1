package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	multiple := 1.0
	for i := 1; i <= n; i++ {
		multiple *= 1.0 + (float64(i) / 10.0)

	}
	fmt.Println(multiple)
}
