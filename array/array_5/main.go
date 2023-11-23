package main

import "fmt"

func main() {
	var n, minOdd int
	var found bool
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if num%2 != 0 && (num < minOdd || !found) {
			minOdd = num
			found = true
		}
	}
	if found {
		fmt.Println(minOdd)
	} else {
		fmt.Println(0)
	}
}
