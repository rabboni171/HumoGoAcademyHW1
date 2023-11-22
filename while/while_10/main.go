package main

import "fmt"

func main() {
	var n, k, powerOfThree int
	fmt.Scan(&n)
	powerOfThree = 1
	k = 0

	for powerOfThree < n {
		powerOfThree *= 3
		k++
	}
	fmt.Println(k - 1)
}
