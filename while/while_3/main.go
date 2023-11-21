package main

import "fmt"

func main() {
	var n, k, rmd, chst int
	fmt.Scan(&n, &k)
	chst = 0 // chastnoe
	rmd = n  // ostatok
	for rmd >= k {
		chst++
		rmd -= k
	}
	fmt.Printf("Chastnoe: %d, Ostatok: %d\n", chst, rmd)
}
