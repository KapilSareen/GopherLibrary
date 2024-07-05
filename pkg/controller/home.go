package controller

import (
	"net/http"

	"github.com/KapilSareen/go-project/pkg/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	views.Home(w, r)
}
