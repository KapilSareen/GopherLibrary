package views

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signup" {
		http.ServeFile(w, r, "templates/signup.html")
	}
	if r.URL.Path == "/login" {
		http.ServeFile(w, r, "templates/login.html")
	}

}
