package main

import "fmt"

func main() {
	var numb int

	fmt.Print("Введите двузначное число: ")

	fmt.Scan(&numb)

	leftDigit := numb / 10
	rightDigit := numb % 10

	fmt.Printf("Левая цифра числа: %d\n", leftDigit)
	fmt.Printf("Правая цифра числа: %d", rightDigit)
}
