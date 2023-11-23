package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Элементы массива", arr)
}
