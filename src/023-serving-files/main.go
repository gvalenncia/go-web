package main

import (
	"net/http"
	"os"
	"log"
	"io"
)

func main ()  {

	//Simple content server
	http.HandleFunc("/dog-content", dogContentHandler)
	//Serving content using ServeFile
	http.HandleFunc("/dog-file", dogFileHandler)
	//Exposing a folder with all assets to serve
	http.HandleFunc("/dog", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dogContentHandler(w http.ResponseWriter, req *http.Request)  {
	f, err := os.Open("dog.jpg")
	defer f.Close()
	handleError(w, err)

	fi, err := f.Stat()
	handleError(w, err)

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

func dogFileHandler(w http.ResponseWriter, req *http.Request)  {
	http.ServeFile(w, req, "assets/dog.jpg")
}

func dog (w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/dog.jpg">`)
}
func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatal("There was an error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
