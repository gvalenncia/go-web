package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

//struct fields must be capitalized, so that the template can see them
//when rendering
type person struct {
	Fname string
	Lname string
}

type dog struct {
	Race string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	//Rendering an slice
	names := []string{"german", "jesus", "ruth"}
	err := tpl.ExecuteTemplate(os.Stdout, "slide.gohtml", names)
	if err != nil {
		log.Fatalln("There was a problem when executing slide.gohtml.", err)
	}

	//Rendering a map
	numbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "maps.gohtml", numbers)
	if err != nil {
		log.Fatalln("There was an error when executing maps.gohtml", err)
	}

	//rendering a struct
	p1 := person{
		Fname: "german",
		Lname: "valencia",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", p1)
	if err != nil {
		log.Fatalln("There was an error when executing struct.gohtml", err)
	}

	//rendering a struct slide
	p2 := person{
		Fname: "Ruth",
		Lname: "Vargas",
	}
	persons := []person{p1, p2}
	err = tpl.ExecuteTemplate(os.Stdout, "struct-slide.gohtml", persons)
	if err != nil {
		log.Fatalln("There was a problem when executing struct-slide.gohtml", err)
	}

	//rendering a composed struct
	d1 := dog{
		Race: "Pastor",
	}
	d2 := dog{
		Race: "Pincher",
	}
	a :=  struct {
		P []person
		D []dog
	}{
		P: []person{p1, p2},
		D: []dog{d1, d2},
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct-composited.gohtml", a)
	if err != nil {
		log.Fatalln("there was a problem when executing struct-composite.gohtml", err)
	}
}
