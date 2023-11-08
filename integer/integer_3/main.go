package main

import "fmt"

func main() {
	var data int // в байтах

	fmt.Println("Введите размер файла:")

	fmt.Scan(&data)

	convertedData := data / 1024

	fmt.Printf("Размер файла в кБ: %d", convertedData)
}
