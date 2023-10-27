package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	os.Stdout.WriteString(strconv.Itoa(sum))
}
