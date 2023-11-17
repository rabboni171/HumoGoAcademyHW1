package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&k, &n)
	for i := 0; i < n; i++ {
		fmt.Println(k)
	}

}
