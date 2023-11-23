package main

import "fmt"

func main() {
	var sizeOfArray int
	fmt.Scan(&sizeOfArray)

	arr := make([]int, sizeOfArray)

	// Заполнение массива
	for i := 0; i < sizeOfArray; i++ {
		fmt.Scan(&arr[i])
	}

	/*Подсчет количества элементов, у которых два соседа
	и которые больше обоих своих соседей*/
	count := 0
	for i := 1; i < sizeOfArray-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
