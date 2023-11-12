package main

import "fmt"

func main() {
	var someNumber int

	fmt.Scan(&someNumber)
	if someNumber > 100 {
		firstDigit := someNumber / 100
		LastDigit := someNumber % 10

		if firstDigit > LastDigit {
			fmt.Println(firstDigit)
		} else if firstDigit < LastDigit {
			fmt.Println(LastDigit)
		}
	}
	fmt.Println("Chislo menshe 100")

}
