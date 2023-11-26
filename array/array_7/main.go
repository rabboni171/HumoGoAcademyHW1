package main

import "fmt"

func main() {
	var cycles int
	fmt.Scan(&cycles)
	for q := 0; q < cycles; q++ {
		var sizeOfArray int
		fmt.Scan(&sizeOfArray)
		arr := make([]int, sizeOfArray)

		// заполняем массив
		for i := 0; i < sizeOfArray; i++ {
			fmt.Scan(&arr[i])
		}

		minElement := arr[0]
		var imin = 0 // индекс минимального элемента
		// берем опорного элемента
		// Пройдем по массиву и найдем min элемент
		for i := 1; i < sizeOfArray; i++ {
			if arr[i] < minElement {
				minElement = arr[i]
				imin = i
			}
			arr[imin]++
		}
		res := 1
		for i := 0; i < sizeOfArray; i++ {
			res *= arr[i]
		}
		fmt.Println(res)

	}

}

// package main

// import "fmt"

// func main() {

//   var t int

//   fmt.Scan(&t)

//   for q := 0; q < t; q++ {
//     var n int

//     fmt.Scan(&n)

//     arr := make([]int, n)

//     for i := 0; i < n; i++ {
//       fmt.Scan(&arr[i])
//     }

//     var minn = arr[0]
//     var iminn = 0
//     for i := 0; i < n; i++ {
//       if minn > arr[i] {
//         minn = arr[i]
//         iminn = i
//       }
//     }

//     arr[iminn]++

//     var res int = 1
//     for i := 0; i < n; i++ {
//       res *= arr[i]
//     }

//     fmt.Println(res)
//   }
// }

/*

5
1 8 2 4 -100
minn = -100
iminn = 4


arr[iminn] ++

*/
