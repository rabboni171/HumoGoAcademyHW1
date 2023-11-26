package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&x)
	for i, num := range arr {
		if num == x {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
