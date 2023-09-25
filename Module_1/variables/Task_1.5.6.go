package main

//Часовая стрелка повернулась с начала суток на d градусов.
//Определите, сколько сейчас целых часов h и целых минут m.

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	hour := d / 30
	min := (d % 30) * 2
	fmt.Print("It is ", hour, " hours ", min, " minutes.")
}
