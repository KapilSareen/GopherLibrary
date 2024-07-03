package controller

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/views"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var Store *sessions.CookieStore

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		views.Login(w, r)
		return
	}

	models.Connect()

	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
    
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	action := r.FormValue("action")
	name := r.FormValue("name")
	password := r.FormValue("password")

	switch action {
	case "Signup":
		err := models.Signup(name, password)
		if err != nil {
			fmt.Fprint(w, "Username already exists. Please choose a different one.")
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	case "Login as Admin":
		user, err := models.Login(name, password)
		if err != nil {
			fmt.Fprint(w, "Incorrect name or password")
			return
		}

		isAdmin := user.IsAdmin
		if !isAdmin {
			fmt.Fprint(w, "You are not an admin. Please log in as a student.")
			return
		}

		session, err := Store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user"] = user.Name
		session.Values["password"] = user.Password
		session.Values["isAdmin"] = user.IsAdmin
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	case "Login as Student":
		user, err := models.Login(name, password)
		if err != nil {
			fmt.Fprint(w, "Incorrect name or password")
			return
		}

		session, err := Store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user"] = user.Name
		session.Values["password"] = user.Password
		session.Values["isAdmin"] = user.IsAdmin

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Print(err.Error())
			return
		}
		http.Redirect(w, r, "/books", http.StatusSeeOther)

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
