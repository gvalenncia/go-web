package main

import (
	"html/template"
	"net/http"
	"log"
	"net/url"
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

	data := struct {
		Method string
		URL *url.URL
		Submissions map[string][]string
		Header http.Header
		Host string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "form.gohtml", data)
}