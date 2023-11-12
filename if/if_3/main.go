package main

import "fmt"

func main() {
	var someNumber int
	fmt.Print("Vvedite eto chislo:")
	fmt.Scan(&someNumber)
	if someNumber > 0 {
		someNumber += 1
	} else if someNumber < 0 {
		someNumber -= 2
	} else {
		someNumber = 10
	}

	fmt.Println("Resultat:", someNumber)

}
