package main

import (
	"net/http"
)

func isLogged(req *http.Request) bool  {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
