package main

// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

// Входные данные
// Программа получает на вход два числа. Гарантируется,
// что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

// Выходные данные
// Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
// Цифры выводятся в порядке их нахождения в первом числе.

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var num1, num2, tmp1, tmp2 int
// 	var str string
// 	fmt.Scan(&num1, &num2)

// 	for num1 > 0 {
// 		tmp1 = num1 % 10
// 		num1 /= 10 // the same: num1 = num1 / 10
// 		tmp2 = num2
// 		for tmp2 > 0 {
// 			if tmp2%10 == tmp1 {
// 				str = strconv.Itoa(tmp1) + " " + str
// 			}
// 			tmp2 /= 10
// 		}
// 	}
// 	fmt.Println(str)
// }

import "fmt"

func main() {
	var num1, num2 string
	//var str string
	fmt.Scan(&num1, &num2)

	for _, a1 := range num1 {
		for _, b1 := range num2 {
			if a1 == b1 {
				fmt.Print(string(a1) + " ")
			}
		}
	}
}
