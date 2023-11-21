package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	isPowerOfThree := false

	for number > 1 {
		if number%3 != 0 {
			isPowerOfThree = false
			break
		}
		number /= 3
		isPowerOfThree = true
	}

	fmt.Println(isPowerOfThree)
}
