package main

import "fmt"

func main() {
	var sizeOfArray int
	fmt.Scan(&sizeOfArray)
	arr := make([]int, sizeOfArray)

	// заполняем массив
	for i := 0; i < sizeOfArray; i++ {
		fmt.Scan(&arr[i])
	}

	maxElement := arr[0] // берем опорного элемента
	// Пройдем по массиву и найдем max элемент
	for i := 1; i < sizeOfArray; i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
	}
	fmt.Println(maxElement)
}
