package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var x int
	fmt.Scan(&x)
	closestElement := arr[0]
	// идея через разность значений по модулю
	minDiff := int(math.Abs(float64(x - arr[0])))
	// пробежимся по ренжу для массива arr
	for _, num := range arr {
		diff := int(math.Abs(float64(x - num)))
		if diff < minDiff {
			minDiff = diff
			closestElement = num
		}
	}
	fmt.Println(closestElement)
}
