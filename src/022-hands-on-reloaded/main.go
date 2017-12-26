package main

import (
	"net"
	"log"
	"io"
)

func main()  {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("there was a problem when listening", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("There was a problem accepting connections", err)
		}

		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
