package controller

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/views"
	"net/http"
)

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	models.Connect()
	p := views.List(w, r)
	Books, err := models.ListBooks()
	if err != nil {
		fmt.Print(err.Error())
	}

	p.Execute(w, Books)
}

func FindBook(w http.ResponseWriter, r *http.Request) {
	g := r.PathValue("book")
	fmt.Printf("Finding book: %s\n", g)
	models.Connect()
	p := views.List(w, r)
	Books, err := models.SearchBook(g)
	if err != nil {
		fmt.Print(err.Error())
	}

	p.Execute(w, Books)
}
