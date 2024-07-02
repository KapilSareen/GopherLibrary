package api

import (
	"github.com/KapilSareen/go-project/pkg/controller"
	"net/http"
	"github.com/KapilSareen/go-project/pkg/middleware"

)

func Start() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.Home)
	mux.HandleFunc("GET /books", middleware.Auth(controller.ShowBooks))
	mux.HandleFunc("GET /book/{book}", middleware.Auth(controller.FindBook))
	mux.HandleFunc("GET /login", controller.Login)
	mux.HandleFunc("POST /login", controller.Login)
	mux.HandleFunc("GET /signup", controller.Login)
	mux.HandleFunc("POST /signup", controller.Login)

	http.ListenAndServe(":8001", mux)
}
