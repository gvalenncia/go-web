package main

import (
	"html/template"
	"net/http"
	"log"
)

type handler int

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	var h handler
	http.ListenAndServe(":8080", h)
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln("There was a problem when parsing the form", err)
	}

	tpl.ExecuteTemplate(w, "form.gohtml", req.Form)
}