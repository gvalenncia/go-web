package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
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
		 go serve(conn)
	}
}

func serve (conn net.Conn)  {
	defer conn.Close()

	sc := bufio.NewScanner(conn)
	for i :=0 ; sc.Scan(); i++ {
		ln := sc.Text()

		if i == 0 {
			fields := strings.Fields(ln)
			fmt.Print(fields[0])
			fmt.Print(fields[1])
		}

		if ln == "" {
			break
		}
	}

	body := "<h1>HOLY COW THIS IS LOW LEVEL</h1>"
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}