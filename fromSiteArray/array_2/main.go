package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var x int
	fmt.Scan(&x)
	var isInArray bool
	for i := 0; i < n; i++ {
		if x == arr[i] {
			isInArray = true
		}
	}
	if isInArray {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
