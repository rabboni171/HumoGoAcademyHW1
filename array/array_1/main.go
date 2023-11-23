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
	hasSameSign := false
	// операции с элементами массива
	for i := 1; i < sizeOfArray; i++ {
		if arr[i]*arr[i-1] > 0 {
			hasSameSign = true
			break
		}
	}
	if hasSameSign {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
