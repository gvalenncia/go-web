package main

import (
	"net/http"
	"html/template"
	json2 "encoding/json"
	"log"
)

var tpl *template.Template

type person struct {
	Fname string
	Lname string
	Items []string
}

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main ()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/mar", mar)
	http.HandleFunc("/enc", enc)
	http.ListenAndServe(":8080", nil)
}
//SAme thing like marshaling but using an object before writing to a response
func enc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"German",
		"Valencia",
		[]string{"guns", "and", "roses"},
	}

	err := json2.NewEncoder(w).Encode(p1)
	if err != nil{
		log.Println("There was en error envoding the json")
	}
}

func mar(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"German",
		"Valencia",
		[]string{"guns", "and", "roses"},
	}

	json, err := json2.Marshal(p1)
	if err != nil {
		log.Println("There was an error marshaling the object")
	}

	w.Write(json)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}