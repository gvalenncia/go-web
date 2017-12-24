package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("There was a problem when running the server", err)
	}
	defer li.Close()

	for{
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("There was a problem accepting connections", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//Killing the connection after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("There was aproblem when setting deadline")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		//Over the connection, do write
		fmt.Fprintf(conn, "Hi German, Do not call me more and go to shit")
	}
	defer conn.Close()
}
