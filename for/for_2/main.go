package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	count := 0
	for i := a; i <= b; i++ {
		fmt.Println(i)
		count++

	}
	fmt.Println(count)

}
