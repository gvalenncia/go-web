package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init ()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main ()  {

	http.HandleFunc("/", index)
	http.HandleFunc("/attend", attend)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Print("In index: ", r.Method)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
	//w.Header().Set("Location", "/")
	//w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func attend(w http.ResponseWriter, r *http.Request) {
	fmt.Print("In attend: ", r.Method)
}

func handleError(w http.ResponseWriter, err error)  {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
