package main

import "fmt"

func main() {
	index := 0
	sum := 0
	var avg float64
	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		index++
		sum += num
		avg = float64(sum) / float64(index)
	}

	fmt.Printf("Avarage result: %0.3f", avg)
}
