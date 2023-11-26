package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	reverseNewArray := reverseArray(arr, n)
	fmt.Println(reverseNewArray)
}

func reverseArray(arr []int, size int) []int {
	lengthArray := len(arr)
	/*
	for i:=0; i < n/2; i++{
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	*/
	for i := 0; i < lengthArray; i++ {
		j := lengthArray - i -1
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]

		} else {
			break
		}
	}
	return arr
}

// for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

// }
