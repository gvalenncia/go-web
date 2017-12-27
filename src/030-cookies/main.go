package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "my value",
	})
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, "my cookie:", c)
}

