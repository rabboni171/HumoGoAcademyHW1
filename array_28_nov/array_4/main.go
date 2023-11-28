package main

import "fmt"

func main() {
	n := readNumber()
	arr := readArray(n)
	result := findNonSquares(arr)
	printResult(result)
}

func readNumber() int {
	var n int
	fmt.Scan(&n)
	return n
}

func readArray(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	return arr
}

// нахождение в массиве чисел, которые не являются квадратами
func findNonSquares(arr []int) []int {
	result := make([]int, 0)
	for _, num := range arr {
		if !isSquare(num, arr) {
			result = append(result, num)
		}
	}
	return result
}

// проверка, является ли числа хорошим квадратом, перебирая элементы массива.
func isSquare(num int, arr []int) bool {
	for _, someNum := range arr {
		if num*num == someNum {
			return true
		}
	}
	return false
}

func printResult(result []int) {
	for i, num := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	//fmt.Println()
}
