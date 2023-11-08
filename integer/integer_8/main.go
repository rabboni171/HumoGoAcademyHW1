package main

import "fmt"

func main() {
	var numb, newNumb int

	fmt.Print("Введите двузначное число: ")

	fmt.Scan(&numb)

	leftDigit := numb / 10
	rightDigit := numb % 10

	newNumb = rightDigit*10 + leftDigit

	fmt.Printf("Число с перестанкой цифр: %d", newNumb)

}
