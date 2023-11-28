package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	printElemsByLines(numbers, k)

}

func printElemsByLines(arr []int, k int) {
	for i := 0; i < len(arr); i += k {
		end := i + k
		if end > len(arr) {
			end = len(arr)
		}
		fmt.Println(arr[i:end])
	}
}
