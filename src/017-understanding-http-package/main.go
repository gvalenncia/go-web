package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
	"os"
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
	defer conn.Close()

	request(conn)

}

func request(conn net.Conn)  {

	scanner := bufio.NewScanner(conn)
	for i := 0; scanner.Scan(); i++  {
		ln := scanner.Text()
		fields := strings.Fields(ln)
		if i == 0 {
			fmt.Fprintln(os.Stdout, "Method: ", fields[0])
			fmt.Fprintln(os.Stdout, "URL: ", fields[1])
		}
	}
}
