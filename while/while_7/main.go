package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	k := 1

	for k*k <= n {
		k++
	}

	fmt.Print(k)
}
