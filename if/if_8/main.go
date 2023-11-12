package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Vvedite eti chisla:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Printf("Max : %.0f, Min : %.0f", a, b)
	} else if b > a {
		fmt.Printf("Max : %.0f, Min : %.0f", b, a)
	} else {
		fmt.Println("a=b")
	}
}
