package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	minGrade := arr[0]
	maxGrade := arr[0]

	// находим мин и макс оценки
	for _, num := range arr {
		if num < minGrade {
			minGrade = num
		}
		if num > maxGrade {
			maxGrade = num
		}
	}
	// Заменяем минимальные на максимальные и наоборот
	for i, grade := range arr {
		if grade == minGrade {
			arr[i] = maxGrade
		} else if grade == maxGrade {
			arr[i] = minGrade
		}
	}
	// Выводим исправленные оценки
	for _, grade := range arr {
		fmt.Print(grade, " ")
	}
}
