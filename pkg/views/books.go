package views

import (
	"fmt"
	"html/template"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) *template.Template {
	page, err := template.ParseFiles("templates/books.html")
	if err != nil {
		fmt.Print(err)
	}
	return page
}
