package main

import "fmt"

func main() {

	var num, num1, num2, num3 int
	fmt.Scan(&num)
	num1 = num / 100
	num2 = num % 100 / 10
	num3 = num % 10

	fmt.Println(num, num1, num2, num3)

}
