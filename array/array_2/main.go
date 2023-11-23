package main

import "fmt"

func main() {
	var sizeOfArray int
	fmt.Scan(&sizeOfArray)
	arr := make([]int, sizeOfArray)
	// заполнение массива
	for i := 0; i < sizeOfArray; i++ {
		fmt.Scan(&arr[i])
	}
	// Вывод элементов массива в обр. порядке
	for i := sizeOfArray - 1; i >= 0; i-- {
		fmt.Printf("%d ",arr[i])
	}
}
