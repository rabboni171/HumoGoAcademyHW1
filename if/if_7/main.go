package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Vvedite chisla a i b:")
	fmt.Scan(&a, &b)

	if a < b {
		fmt.Println("Poryadkovii nomer menshego chisla :1")
	} else if a > b {
		fmt.Println("Poryadkovii nomer menshego chisla :2")
	} else {
		fmt.Println("Chisla ravni")
	}
}
