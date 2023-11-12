package main

import "fmt"

func main() {
	var someNumber int

	fmt.Print("Vvedite chislo:")

	fmt.Scan(&someNumber)

	fmt.Println((someNumber % 2) == 0)
}
