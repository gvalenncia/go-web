package main

import (
	"html/template"
	"net/http"
	"log"
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
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/submit-signup", submitSignup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, req *http.Request) {
	if !isLogged(req) {
		fmt.Println("User is not logged")
		http.Redirect(w,req,"/signup", http.StatusSeeOther)
		return
	}

	c, err := req.Cookie("session")
	if err != nil {
		log.Println("User not logged in")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un := dbSessions[c.Value]
	u := dbUsers[un]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}


func submitSignup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("first")
		l := req.FormValue("last")

		newu := user{
			UserName: un,
			First: f,
			Last: l,
		}

		if _, ok := dbUsers[un]; ok {
			log.Fatal("username already taken")
			http.Error(w, "username already taken", http.StatusForbidden)
			return
		}

		sessionID := uuid.NewV4().String()
		c := &http.Cookie{
			Name: "session",
			Value: sessionID,
		}

		dbSessions[sessionID] = un
		dbUsers[un] = newu

		http.SetCookie(w, c)
		fmt.Println("user created and logged")
		http.Redirect(w, req, "/bar", http.StatusSeeOther)
	}
}