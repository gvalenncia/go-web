package main

import (
	"net/http"
	"github.com/satori/go.uuid"
)

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {

	scookie, err := r.Cookie("session")

	if err != nil {
		id := uuid.Must(uuid.NewV4())
		scookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
		}

		http.SetCookie(w, scookie)
	}
}