package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Germán")
	if err != nil {
		log.Fatalln("There was a problem when executing idex.gohtml", err)
	}
}
