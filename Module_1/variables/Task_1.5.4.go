package main

//Дано натуральное число, выведите его последнюю цифру.

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	res := num % 10
	fmt.Println(res)
}
