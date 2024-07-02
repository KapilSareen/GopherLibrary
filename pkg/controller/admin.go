package controller

import (
	"net/http"
	"github.com/KapilSareen/go-project/pkg/views"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}
	views.Admin(w, r)

}
