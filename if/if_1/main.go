package main

import "fmt"

func main() {
	var someNumber int
	fmt.Print("Vvedite eto chislo:")
	fmt.Scan(&someNumber)
	if someNumber > 0 {
		fmt.Println(someNumber + 1)
	} else {
		fmt.Println(someNumber)
	}

}
