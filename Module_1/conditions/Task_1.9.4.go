package main

// Определите является ли билет счастливым. Счастливым считается билет,
// в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

import "fmt"

func main() {

	var num uint
	fmt.Scan(&num)

	num0 := num / 100000
	num1 := num % 100000 / 10000
	num2 := num % 10000 / 1000
	num3 := num % 1000 / 100
	num4 := num % 100 / 10
	num5 := num % 10

	if (num0 + num1 + num2) == (num3 + num4 + num5) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
