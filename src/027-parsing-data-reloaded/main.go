package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"fmt"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init ()  {
	tpl = template.Must(template.ParseFiles("form.gohtml"))
}

func main () {
	http.HandleFunc("/", index)
	http.HandleFunc("/attend", formHandler)
	http.Handle("/favicon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	f,h,err := r.FormFile("qfile")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//reading a file and sending back the content
	bs, err := ioutil.ReadAll(f)

	fmt.Print(string(bs))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tpl.Execute(w, string(bs))

	//save a file in the disk
	dst, err := os.Create(filepath.Join(".", h.Filename + "_new"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer dst.Close()

	_, err = dst.Write(bs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request)  {
	tpl.Execute(w, nil)
}
