package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", nil)

	if err != nil {
		log.Fatalln("there was an error when parsing tpl1.gohtml", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gmao", nil)

	if err != nil {
		log.Fatalln("there was an error when parsing tpl2.gmao", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", nil)

	if err != nil {
		log.Fatalln("there was an error when parsing tpl3.gohtml", err)
	}
}
