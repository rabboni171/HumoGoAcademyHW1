package main

import "fmt"

func main() {
	var a int
	// cnt = 0
	num := 0
	for {
		fmt.Scan(&a)
		if a != 0 {
			if a > num {
				num = a
			}
			continue
		} else {
			fmt.Println(num)
			break
		}

	}
}
