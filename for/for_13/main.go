package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0.0
	index := 1.0
	for i := 1; i <= n; i++ {
		sum += float64(i) * index
		index = -index

	}
	fmt.Println(sum)
}
