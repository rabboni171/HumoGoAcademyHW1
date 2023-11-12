package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Vvedite eti chisla:")
	fmt.Scan(&a, &b)

	if a != b {
		sum := a + b
		a += sum
		b += sum
	} else if a == b {
		a, b = 0, 0
	}
	fmt.Println("Novoe znachenie a:", a)
	fmt.Println("Novoe znachenie b:", b)

}
