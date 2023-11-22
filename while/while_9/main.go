package main

import "fmt"

func main() {
	var n, k, powerOfThree int
	fmt.Scan(&n)
	powerOfThree = 1
	for powerOfThree <= n {
		k++
		powerOfThree *= 3
	}
	fmt.Println(k)
}
