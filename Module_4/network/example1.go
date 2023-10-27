package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	message := make([]byte, 1024) // создадим буфер
	n, err := conn.Read(message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(message[:n])) // напечатаем сообщение

}
