package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/req", req)
	http.ListenAndServe(":8080", nil)
}

func req(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "response from req")
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml",nil)
}
