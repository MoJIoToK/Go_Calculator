package main

// Последовательность состоит из натуральных чисел и завершается числом 0.
// Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

import "fmt"

func main() {
	var num, res, max int
	res = 1
	max = 0
	for fmt.Scan(&num); num != 0; fmt.Scan(&num) {
		if num > max {
			max = num
			res = 1
		} else if max == num {
			res++
		}
	}
	fmt.Println(res)
}
