package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Vvedite chisla a i b:")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Max:", a)
	} else {
		fmt.Println("Max:", b)
	}

}
