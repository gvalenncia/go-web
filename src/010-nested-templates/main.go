package main

import (
	"go/ast"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "German")
	if err != nil {
		log.Fatalln("There was a problem when executing idex.gohtml", err)
	}

}
