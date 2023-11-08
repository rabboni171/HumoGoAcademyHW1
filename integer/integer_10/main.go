package main

import "fmt"

func main() {
	var numb, lastDigit, secondDigit int

	fmt.Print("Введите трехзначное число: ")

	fmt.Scan(&numb)

	lastDigit = numb % 10

	secondDigit = (numb / 10) % 10

	fmt.Printf("Последняя цифра числа: %d\n", lastDigit)

	fmt.Printf("Средняя цифра числа: %d", secondDigit)

}
