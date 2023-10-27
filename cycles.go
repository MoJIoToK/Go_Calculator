package main

import "fmt"

func main() {
	var num int
	for fmt.Scan(&num); num < 101; fmt.Scan(&num) {
		if num >= 10 {
			fmt.Println(num)
		} else if num < 10 {
			continue
		}
	}
}
