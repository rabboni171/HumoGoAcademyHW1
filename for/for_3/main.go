package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	count := 0
	for i := b - 1; i > a; i-- {
		fmt.Println(i)
		count++

	}
	fmt.Println(count)

}
