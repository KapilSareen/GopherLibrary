package controller
import(
"fmt"
"time"
"github.com/KapilSareen/go-project/pkg/models"
"github.com/KapilSareen/go-project/pkg/types"

)
func AddBook(){
models.Connect()
book:=types.Book{
	Name:      "TEST-add",
	Author:    "TEST-auth",
	Price:     69.69,
	OwnedFrom: time.Now().String(),
	IsAvail:   true,     
}
fmt.Print(book)
models.AddBook(book)
}