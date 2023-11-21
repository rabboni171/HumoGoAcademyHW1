package main

import "fmt"

func main() {
	sum := 0
	num := 0

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
	}

	fmt.Println(sum)
}
