package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	k := 0

	for k*k <= n {
		k++
	}
	k--
	fmt.Print(k)
}
