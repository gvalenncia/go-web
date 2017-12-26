package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("There was a problem when running the server", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("There was a problem accepting connections", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	handleRequest(conn)
}

func handleRequest(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for i := 0; scanner.Scan(); i++ {
		var ln string = scanner.Text()
		if i == 0 {
			fields := strings.Fields(ln)
			httpMux(conn, fields[0], fields[1])
		} else {
			break
		}
	}
}

func httpMux(conn net.Conn, method string, uri string) {

	if method == "GET" && uri == "/" {
		index(conn)
	} else if method == "GET" && uri == "/home" {
		index(conn)
	} else if method == "GET" && uri == "/about" {
		about(conn)
	} else if method == "GET" && uri == "/contact" {
		contact(conn)
	} else {
		notFound(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charet="UTF-8">
				<title></title>
				</head>
					<body>
						<strong>INDEX</strong><br>
						<a href="/">index</a><br>
						<a href="/about">about</a><br>
						<a href="/contact">contact</a><br>
					</body>
				</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charet="UTF-8">
				<title></title>
				</head>
					<body>
						<strong>About me</strong><br>
						<a href="/">index</a><br>
						<p>This is the about me page</p>
					</body>
				</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charet="UTF-8">
				<title></title>
				</head>
					<body>
						<strong>Contact me</strong><br>
						<a href="/">index</a><br>
						<p>Reach me at: gvalenncia@gmail.com</p>
					</body>
				</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func notFound(conn net.Conn) {
	body := `<!DOCTYPE html>
				<html lang="en">
				<head>
				<meta charet="UTF-8">
				<title></title>
				</head>
					<body>
						<strong>NOT FOUND</strong><br>
						<a href="/">index</a><br>
					</body>
				</html>`
	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
