package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result := 1.0
	if n%2 == 0 {
		for i := n; i >= 2; i -= 2 {
			result *= float64(i)
		}
	} else {
		for i := n; i >= 1; i -= 2 {
			result *= float64(i)
		}
	}

	fmt.Println(result)
}
