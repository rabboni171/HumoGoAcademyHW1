package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Vvedite eti chisla:")
	fmt.Scan(&a, &b, &c)

	count1 := 0
	count2 := 0

	if a > 0 {
		count1++
	} else if a < 0 {
		count2++
	}
	if b > 0 {
		count1++
	} else if b < 0 {
		count2++
	}
	if c > 0 {
		count1++
	} else if c < 0 {
		count2++
	}

	fmt.Println("Kol-vo polozhitelnih", count1)
	fmt.Println("Kol-vo otricatelnih", count2)
}
