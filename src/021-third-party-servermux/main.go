package main

import (
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.POST("/send-message", sendMessage)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params)  {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	handleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params)  {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	handleError(w, err)
}

func sendMessage(w http.ResponseWriter, req *http.Request, ps httprouter.Params)  {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	fmt.Print(ps)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, e error) {
	if e != nil {
		http.Error(w, e.Error() ,http.StatusInternalServerError)
		log.Fatal("There was a problem when executing template: ", e)
	}
}

