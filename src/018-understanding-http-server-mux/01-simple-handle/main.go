package _1_simple_handle

import (
	"net/http"
	"io"
)

type homeHandler int
type aboutHandler int

func main()  {
	var homeh homeHandler
	var abouth aboutHandler

	mux := http.NewServeMux()
	//It will catch everything behind home/*
	mux.Handle("/home/", homeh)
	mux.Handle("/about", abouth)

	http.ListenAndServe(":8080", mux)
}

func (hh homeHandler) ServeHTTP (resw http.ResponseWriter, req *http.Request)  {
	io.WriteString(resw, "<h1>From Home</h1>")
}

func (ah aboutHandler) ServeHTTP (resw http.ResponseWriter, req *http.Request)  {
	io.WriteString(resw, "<h1>From About</h1>")
}
