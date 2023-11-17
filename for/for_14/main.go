package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (2*i - 1)
		// sum = sum + (2*i-1)
		fmt.Println(sum)
	}
}
