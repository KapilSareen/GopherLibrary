package controller

import (
	"fmt"
	"net/http"
	// "github.com/KapilSareen/go-project/pkg/models"

)

func Home(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/" {
	http.NotFound(w, r)
	return
}
// models.Connect()

// a,err:=models.SearchBook("TEST-author")

// if(err!=nil){
	
// 	fmt.Print(err)
// }
fmt.Fprint(w,"<h1>Welcome page</h1>")
// views.Home(w , r )




}	