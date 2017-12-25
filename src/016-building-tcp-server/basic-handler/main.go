package main

import (
	"net/http"
	"fmt"
)

type handler int

func main()  {
	var h handler
	http.ListenAndServe(":8080", h)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>This is a handler</h1>")
}