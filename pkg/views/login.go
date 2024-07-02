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


func Admin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/admin" {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
        w.Header().Set("Pragma", "no-cache")
        w.Header().Set("Expires", "0")
		http.ServeFile(w, r, "templates/admin.html")
	}
}