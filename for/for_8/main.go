package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	multiple := 1
	for i := a; i <= b; i++ {
		multiple = multiple * i

	}
	fmt.Println(multiple)

}
