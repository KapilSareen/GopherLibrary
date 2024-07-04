package api

import (
	"github.com/KapilSareen/go-project/pkg/controller"
	"github.com/KapilSareen/go-project/pkg/middleware"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Home)

	mux.HandleFunc("GET /books", middleware.Auth(controller.ShowBooks))
	mux.HandleFunc("POST /books", middleware.Auth(controller.AddBook))

	mux.HandleFunc("POST /receive", middleware.Auth(controller.ReceiveBook))
	mux.HandleFunc("POST /issue", middleware.Auth(controller.IssueBook))
	mux.HandleFunc("POST /delete", middleware.Auth(controller.DeleteBook))

	mux.HandleFunc("GET /book/{book}", middleware.Auth(controller.FindBook))

	mux.HandleFunc("GET /login", controller.Login)
	mux.HandleFunc("POST /login", controller.Login)

	mux.HandleFunc("GET /signup", controller.Login)
	mux.HandleFunc("POST /signup", controller.Login)

	mux.HandleFunc("GET /admin", middleware.Auth(middleware.IsAdmin(controller.Admin)))

	mux.HandleFunc("GET /logout", controller.Logout)

	http.ListenAndServe(":8000", mux)
}
