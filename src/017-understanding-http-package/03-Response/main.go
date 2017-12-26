package main

import (
	"net/http"
	"fmt"
)

type handler int

func main() {

	var h handler
	http.ListenAndServe(":8080", h)
}

func (h handler) ServeHTTP(resw http.ResponseWriter, req *http.Request)  {
	resw.Header().Set("test-header", "this is a test header")
	resw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(resw, "<h1>Hello Baby</h1>")
}