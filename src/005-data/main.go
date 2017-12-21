package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("slide.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "slide.gohtml", "german")

	if err != nil {
		log.Fatalln("there was a problem when executing slide.gohtml template", err)
	}
}
