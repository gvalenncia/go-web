package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

type person struct {
	Fname string
	Lname string
	Age string
	Qs string
}

func init ()  {
	tpl = template.Must(template.ParseFiles("person.gohtml"))
}

func main () {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	qs := r.FormValue("qs")

	f := r.FormValue("fname")
	l := r.FormValue("lname")
	a := r.FormValue("age")

	p := person{
		Fname:f,
		Lname:l,
		Age:a,
		Qs:qs,
	}

	tpl.Execute(w, p)
}
