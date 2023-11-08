package main

import "fmt"

func main() {
	var numb, firsDigit int

	fmt.Print("Введите трехзначное число: ")

	fmt.Scan(&numb)

	firsDigit = numb / 100

	fmt.Printf("Первая цифра числа: %d", firsDigit)

}
