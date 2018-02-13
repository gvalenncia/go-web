package main

import (
	"html/template"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

var tpl *template.Template

type code struct {
	Code int
	Description string
}

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	var data []code
	rcvd := `[{"Code":200,"Description":"StatusOK"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":302,"Description":"StatusFound"},{"Code":303,"Description":"StatusSeeOther"},{"Code":307,"Description":"StatusTemporaryRedirect"},{"Code":400,"Description":"StatusBadRequest"},{"Code":401,"Description":"StatusUnauthorized"},{"Code":402,"Description":"StatusPaymentRequired"},{"Code":403,"Description":"StatusForbidden"},{"Code":404,"Description":"StatusNotFound"},{"Code":405,"Description":"StatusMethodNotAllowed"},{"Code":418,"Description":"StatusTeapot"},{"Code":500,"Description":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data{
		fmt.Println(v.Code, v.Description)
	}
}
