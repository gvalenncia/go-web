package main

import (
	"html/template"
	"strings"
	"os"
	"log"
	"time"
	"math"
)

var tpl *template.Template

/*
We can use template.FuncMap interface to define 
the set of functions we want to use in templates
*/
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"t": time.Now,
	"fdbl": double,
	"fsqr": square,
}

func init ()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main ()  {
	data := []string{
		"German",
		"Valencia",
		"Vargas",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "func-map.gohtml", data)
	if err != nil {
		log.Fatalln("There was an error executing func-map.gohtml")
	}
}

func firstThree(s string) string  {
	s = strings.TrimSpace(s)
	return s[:3]
}

func square (x int) int  {
	return x * x
}

func double (x int)  float64 {
	return math.Pow(float64(x), 2)
}