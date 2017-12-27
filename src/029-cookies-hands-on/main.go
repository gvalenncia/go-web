package main

import (
	"net/http"
	"strconv"
	"io"
)

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, request *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "session",
		Value: "my value",
	})
}

func read(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	io.WriteString(w, c.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session")
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusOK)
}


func index(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("my-cookie")

	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name: "my-cookie",
			Value: "0",
		})
	}

	counter,_ := strconv.Atoi(c.Value)
	counter++
	c.Value = strconv.Itoa(counter)

	http.SetCookie(w, c)
	io.WriteString(w, c.Value)
}
