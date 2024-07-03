package controller

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/types"
	"net/http"
	"strconv"
	"time"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	models.Connect()

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := r.FormValue("name")
	author := r.FormValue("author")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 32)
	book := types.Book{
		Name:      name,
		Author:    author,
		Price:     float32(price),
		OwnedFrom: time.Now().String(),
		IsAvail:   true,
	}

	err:=models.AddBook(book)
	if(err!=nil){
		fmt.Fprint(w,"Book already exists")
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
