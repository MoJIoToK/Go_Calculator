package main

//По данному трехзначному числу определите, все ли его цифры различны.

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)
	num1 := num / 100
	num2 := num % 100 / 10
	num3 := num % 10

	if num1 != num2 && num2 != num3 && num1 != num3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
