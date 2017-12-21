package main

import (
	"html/template"
	"os"
	"log"
	"net/http"
	"go/ast"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "German")
	if err != nil {
		log.Fatalln("There was a problem when executing idex.gohtml", err)
	}

}
