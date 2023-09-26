package main

// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
// Программа в первой строке получает на вход число n - количество чисел
// в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

import "fmt"

func main() {
	var size, num, sum, i int
	fmt.Scan(&size)
	sum = 0
	i = 1
	for fmt.Scan(&num); i < size; fmt.Scan(&num) {
		if num%8 == 0 && 9 < num && num < 100 {
			sum += num
		}
		i++
	}
	fmt.Println(sum)
}
