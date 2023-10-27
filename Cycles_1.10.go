package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2, tmp1, tmp2 int
	var str string
	fmt.Scan(&num1, &num2)

	for num1 > 0 {
		tmp1 = num1 % 10
		num1 = num1 / 10
		tmp2 = num2
		for tmp2 > 0 {
			if tmp2%10 == tmp1 {
				str = strconv.Itoa(tmp1) + " " + str
			}
			tmp2 = tmp2 / 10
		}
	}
	fmt.Println(str)
}
