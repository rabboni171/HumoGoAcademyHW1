package main

import "fmt"

func main() {
	var a, b, cnt int
	fmt.Scan(&a, &b)
	cnt = 0
	for a >= b {
		a -= b
		cnt++
	}
	fmt.Println(cnt)
}
