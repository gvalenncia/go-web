package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln("there was a problem when parsing template: ", err)
	}

	nf, err := os.Create("index.gohtml")

	if err != nil {
		log.Fatalln("There was a problem creating index.gohtml", err)
	}

	defer nf.Close()

	tpl.Execute(nf, nil)

}
