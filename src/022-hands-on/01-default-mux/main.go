package main

import (
	"net/http"
	"fmt"
	"html/template"
)

var tpl *template.Template
type doghandler int
type mehandler int

func init ()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {

	var dh doghandler
	var mh mehandler

	mux := http.NewServeMux()
	mux.Handle("/dog", dh)
	mux.Handle("/me/", mh)

	http.ListenAndServe(":8080", mux)
}

func (dh doghandler) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	data := "firulais"
	tpl.Execute(w, data)
}

func (mh mehandler) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(w, "<p>German Valencia</p>")
}