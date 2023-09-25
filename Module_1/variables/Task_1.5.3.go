package main

//По данному целому числу, найдите его квадрат. На вход подаётся одно целое число.

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	res := num * num
	fmt.Println(res)
}
