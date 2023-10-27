package main

import (
	"log"
	"net"
)

func main() {

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()
	// Открываем порт
	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write([]byte("message"))
	if err != nil {
		log.Println(err)
	}

}
