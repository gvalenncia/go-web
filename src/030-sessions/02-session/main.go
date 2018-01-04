package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

var tpl *template.Template

type user struct {
	UserName string
	First string
	Last string
}

var dbUsers = map[string]user{} //user ID, user
var dbSessions = map[string]string{} //session ID, user ID

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/getuser", getUser)
	http.ListenAndServe(":8080", nil)
}

func getUser(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("no cookie")
		http.Redirect(w, req,"/", http.StatusSeeOther)
		return
	}

	username, ok := dbSessions[c.Value]

	if !ok {
		fmt.Println("no user associated to this session")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[username]

	tpl.ExecuteTemplate(w, "user.gohtml", u)
}

func index(w http.ResponseWriter, req *http.Request) {

	//Create a cookie if it does not exist
	c, err := req.Cookie("sesssion")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w,c)
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("first")
		l := req.FormValue("last")

		u := user{un,f,l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
