package main

//Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

import "fmt"

func main() {
	var num, res int
	fmt.Scan(&num)
	res = num % 100 / 10
	fmt.Println(res)
}
