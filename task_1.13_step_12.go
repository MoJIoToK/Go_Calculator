package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Scan(&num)

	for res := 1; res < num; {
		res = res * 2
		fmt.Println(res)
	}

}
