package main

import "fmt"

func main() {
	index := 0

	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num%2 == 0 {
			index++
		}

	}

	fmt.Println(index)
}

// package main

// import "fmt"

// func main() {
// 	var a, cnt, zeroCount int64
// 	cnt = 0
// 	zeroCount = 0
// 	for {
// 		fmt.Scan(&a)
// 		if a != 0 {
// 			cnt++
// 			zeroCount = 0
// 		} else {
// 			zeroCount++
// 			if zeroCount == 2 {
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(cnt)
// }
