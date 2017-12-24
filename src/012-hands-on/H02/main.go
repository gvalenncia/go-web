package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type table struct {
	Date string
	Open string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	cvsFile := "table.csv"
	file, err := os.Open(cvsFile)
	defer file.Close()
	if err != nil {
		log.Fatalln("There was a problem when opening the file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tables []table

	for i := 0; scanner.Scan(); i++ {
		if i != 0 {
			eSlide := strings.Split(scanner.Text(), ",")
			t := table{
				Date: eSlide[0],
				Open: eSlide[1],
			}
			tables = append(tables, t)
		}
	}

	fmt.Println(tables)
	tpl.Execute(os.Stdout, tables)
}
