package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init ()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main()  {

	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

