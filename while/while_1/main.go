package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	rmd := a
	for rmd >= b {
		rmd -= b
	}
	fmt.Println(rmd)
}
