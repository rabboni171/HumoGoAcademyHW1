package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Vvedite cherez probel chisla a, b, c:")
	fmt.Scan(&a, &b, &c)
	fmt.Println(b < a && b < c)
}
