package main

import (
	"net/http"
	"fmt"
	"html/template"
)

var tpl *template.Template

func init ()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {

	http.HandleFunc("/dog", doghandler)
	http.HandleFunc("/me/", mehandler)
	http.ListenAndServe(":8080", nil)
}

func doghandler (w http.ResponseWriter, req *http.Request)  {
	data := "firulais"
	tpl.Execute(w, data)
}

func mehandler (w http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(w, "<p>German Valencia</p>")
}