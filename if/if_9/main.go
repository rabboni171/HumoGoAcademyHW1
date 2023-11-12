package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Vvedite eti chisla: ")
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}
	fmt.Println("Novoe znachenie a:", a)
	fmt.Println("Novoe znachenie b:", b)

}
