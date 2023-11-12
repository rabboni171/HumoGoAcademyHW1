package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Vvedite cherez probel chisla a i b:")
	fmt.Scan(&a, &b)
	fmt.Println(a > 2 && b <= 3)
}
