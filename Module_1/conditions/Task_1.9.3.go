package main

//Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

import "fmt"

func main() {

	var num uint
	fmt.Scan(&num)
	switch {
	case num <= 9:
		fmt.Println(num)
	case num > 9 && num < 100:
		fmt.Println(num / 10)
	case num > 99 && num < 1000:
		fmt.Println(num / 100)
	case num > 999 && num < 10000:
		fmt.Println(num / 1000)
	default:
		fmt.Println(num / 10000)
	}
}
