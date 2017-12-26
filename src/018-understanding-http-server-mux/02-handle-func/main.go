package main

import (
	"net/http"
	"io"
)

func main()  {

	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler (resw http.ResponseWriter, req *http.Request)  {
	io.WriteString(resw, "<h1>From Home</h1>")
}

func aboutHandler (resw http.ResponseWriter, req *http.Request)  {
	io.WriteString(resw, "<h1>From About</h1>")
}
