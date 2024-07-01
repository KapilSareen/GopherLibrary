package controller

import (
	"fmt"
	"net/http"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/views"
)

func ShowBooks(w http.ResponseWriter, r *http.Request){
	models.Connect()
	p:=views.List(w,r)
	Books,err:=models.ListBooks()
	if(err!=nil){
		fmt.Print(err.Error())
	}
	
	p.Execute(w,Books)
}