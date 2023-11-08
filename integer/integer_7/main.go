package main

import "fmt"

func main() {
	var numb int

	fmt.Print("Введите двузначное число: ")

	fmt.Scan(&numb)

	leftDigit := numb / 10
	rightDigit := numb % 10

	sumOfDigits := leftDigit + rightDigit

	multipleOfDigits := leftDigit * rightDigit

	fmt.Printf("Cумма цифр: %d\n", sumOfDigits)
	fmt.Printf("Произведение цифр: %d", multipleOfDigits)
}
