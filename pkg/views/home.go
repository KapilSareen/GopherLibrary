package views

import (
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/welcome.html")
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/404.html")
}
