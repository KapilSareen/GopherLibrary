package controller

import (
	"net/http"

	"github.com/KapilSareen/go-project/pkg/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        views.ErrorPage(w,r)
		return
	}
	views.Home(w, r)
}
