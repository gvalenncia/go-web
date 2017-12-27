package main

import (
	"net/http"
	"strconv"
	"io"
)

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
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
