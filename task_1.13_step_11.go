package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Scan(&num)

	if num > 10 && num < 15 {
		fmt.Printf("%d korov", num)
	} else {
		tmp := num % 10
		if tmp == 1 {
			fmt.Printf("%d korova", num)
		} else if tmp > 1 && tmp < 5 {
			fmt.Printf("%d korovy", num)
		} else if tmp == 0 || (tmp > 4 && tmp < 9) {
			fmt.Printf("%d korov", num)
		}
	}

}
