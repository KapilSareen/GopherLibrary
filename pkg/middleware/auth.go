package middleware

import (
	"github.com/KapilSareen/go-project/pkg/controller"
	"github.com/KapilSareen/go-project/pkg/models"
	"net/http"
)

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if controller.Store == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		session, err := controller.Store.Get(r, "session")
		if err != nil {
			panic(err)
		}
		name, _ := session.Values["user"].(string)
		if name == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		f(w, r)
	}
}

func IsAdmin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if controller.Store == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		session, err := controller.Store.Get(r, "session")
		if err != nil {
			panic(err)
		}
		name, _ := session.Values["user"].(string)
		models.Connect()
		isAdmin := models.IsAdmin(name)
		if !isAdmin {
			http.Error(w, "You are not authorized", http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}
