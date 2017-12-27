package main

import (
	"html/template"
	"net/http"
	"io"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/doggie", dogImg)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>foo ran</h1>")
}
