package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type course struct {
	Code    string
	Subject string
	Grade   float64
}

type student struct {
	Fname   string
	Lname   string
	Courses []course
}

type school struct {
	Name     string
	Students []student
}

var fm = template.FuncMap{
	"ag": averageGrade,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {

	data := school{
		Name: "inpahu",
		Students: []student{
			{
				Fname: "German",
				Lname: "Valencia",
				Courses: []course{
					{
						Code:    "C-001",
						Subject: "Maths",
						Grade:   4.5,
					},
					{
						Code:    "C-002",
						Subject: "Arts",
						Grade:   3.9,
					},
				},
			},
			{
				Fname: "Jose",
				Lname: "Arias",
				Courses: []course{
					{
						Code:    "C-005",
						Subject: "Social affairs",
						Grade:   4.3,
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

func averageGrade(s student) float64 {
	res := 0.0
	div := len(s.Courses)
	for _, course := range s.Courses {
		res = res + course.Grade
	}

	return res / float64(div)
}
