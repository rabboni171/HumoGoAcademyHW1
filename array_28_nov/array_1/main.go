package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	result := removeElement(numbers)
	fmt.Println(result)
}
func removeElement(arr []int) []int {
	minPosElm := math.MaxInt
	maxNegElm := math.MinInt
	for _, value := range arr {
		if value > 0 && value < minPosElm {
			minPosElm = value
		}
		if value < 0 && value > maxNegElm {
			maxNegElm = value
		}

	}
	// для проверки
	// fmt.Println(minPosElm)
	// fmt.Println(maxNegElm)
	// удаление минимального полож элемента
	for i, value := range arr {
		if value == minPosElm {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	// удаление максимального отриц элемента
	for i, value := range arr {
		if value == maxNegElm {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	return arr
}
