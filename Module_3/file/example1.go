package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	b := bytes.NewReader([]byte("Данные в объекте io.Reader"))
	data, err := ioutil.ReadAll(b)
	if err != nil {
		// ...
	}

	fmt.Printf("%s\n", data) // Данные в объекте io.Reader

}
