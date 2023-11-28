package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	printElements(numbers)

}

func printElements(arr []int) {
	// используем принцип два указателя
	n := len(arr)
	left := 0
	right := n - 1

	// Цикл для итерации по слайсу
	for left <= right {
		// Выводим элемент с начала и, если возможно, элемент с конца
		fmt.Print(arr[left], " ")

		if left != right {
			fmt.Print(arr[right], " ")
		}
		left++
		right--
	}
	fmt.Println()
}
