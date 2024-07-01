package api

import (
	"net/http"
	"github.com/KapilSareen/go-project/pkg/controller"
	
)

 func Start(){
    http.HandleFunc("GET /", controller.Home)
	http.HandleFunc("GET /books", controller.ShowBooks)

	http.ListenAndServe(":8001",nil)
 }