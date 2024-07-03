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
		return
	}

	p.Execute(w, Books)
}

func FindBook(w http.ResponseWriter, r *http.Request) {
	g := r.PathValue("book")
	models.Connect()
	p := views.List(w, r)
	Books, err := models.SearchBook(g)
	if err != nil {
		return
	}

	p.Execute(w, Books)
}

func IssueBook(w http.ResponseWriter, r *http.Request) {

	models.Connect()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	book := r.FormValue("book")
	issued_to := r.FormValue("issued_to")
	UserExist := models.UserExist(issued_to)
	if !UserExist {
		fmt.Fprint(w, "User does not exists")
		return
	}
	models.Connect()
	BookAvail := models.BookAvail(book)
	if !BookAvail {
		fmt.Fprint(w, "Book does not exists or is not available rn")
		return
	}
	models.Connect()
	models.Issue(book, issued_to)
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	models.Connect()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	book := r.FormValue("book")
	BookExist := models.BookExist(book)
	if !BookExist {
		fmt.Fprint(w, "Book does not exists")
		return
	}
	models.Connect()
	models.DeleteBook(book)
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
