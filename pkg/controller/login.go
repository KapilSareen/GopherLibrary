package controller

import (
	"fmt"
	"net/http"
	"os"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/views"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		views.Login(w, r)
		return
	}
	
	models.Connect()
	
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	fmt.Print(Store)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	action := r.FormValue("action")
	name := r.FormValue("name")
	password := r.FormValue("password")

	
	switch action {
	case "Signup":
		fmt.Println("Signup request received")
		err := models.Signup(name, password)
		if err != nil {
			fmt.Fprint(w, "Username already exists. Please choose a different one.")
			return
		}

	case "Login":
		fmt.Println("Login request from student")
		user, err := models.Login(name, password)
		if err != nil {
			fmt.Fprint(w, "Incorrect name or password")
			return
		}

		session, err := Store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Print("Here")
			return
		}

		session.Values["user"] = user.Name
		session.Values["password"]=user.Password
		session.Values["isAdmin"] = user.IsAdmin
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Print(err.Error())
			return
		}

		fmt.Printf("Logged in as %s\n", user.Name)

	case "AdminLogin":
		var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
		fmt.Println("Admin Login request")
		user, err := models.Login(name, password)
		if err != nil {
			fmt.Fprint(w, "Incorrect name or password")
			return
		}

		isAdmin := models.IsAdmin(user)
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
		session.Values["password"]=user.Password
		session.Values["isAdmin"] = user.IsAdmin
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		

		fmt.Printf("Admin logged in as %s\n", user.Name)
	}

	http.Redirect(w,r,"/books",http.StatusSeeOther)
	// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
}
