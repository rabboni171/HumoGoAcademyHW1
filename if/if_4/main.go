package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Vvedite eti chisla:")
	fmt.Scan(&a, &b, &c)

	count := 0

	if a > 0 {
		count++
	}
	if b > 0 {
		count++
	}
	if c > 0 {
		count++
	}

	fmt.Println("Kol-vo", count)
}
