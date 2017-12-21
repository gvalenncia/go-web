package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template


type course struct {
	Code string
	Subject string
}

type student struct{
	Fname string
	Lname string
	Courses []course
}

type school struct {
	Name string
	Students []student
}

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main ()  {

	data := school{
		Name: "inpahu",
		Students: []student{
			{
				Fname: "German",
				Lname: "Valencia",
				Courses: []course{
					{
						Code: "C-001",
						Subject: "Maths",
					},
					{
						Code: "C-002",
						Subject:"Arts",
					},
				},
			},
			{
				Fname: "Jose",
				Lname: "Arias",
				Courses: []course{
					{
						Code: "C-005",
						Subject: "Social affairs",
					},
				},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln("There was a problem when executing index.gohtml", err)
	}
}
