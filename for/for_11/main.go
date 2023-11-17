package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := n; i <= 2*n; i++ {
		sum += i * i

	}
	fmt.Println(sum)

}
