package middleware

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/controller"
	"github.com/KapilSareen/go-project/pkg/models"
	"net/http"
)

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am middleware")

		session, err := controller.Store.Get(r, "session")
		if err != nil {
			panic(err)
		}

		name, _ := session.Values["user"].(string)
		password, _ := session.Values["password"].(string)
		models.Connect()

		user, err := models.Login(name, password)
		if user.Name == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		fmt.Print(user, err)
		f(w, r)

	}
}
