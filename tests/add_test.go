package test

import (
	"fmt"
	"github.com/KapilSareen/go-project/pkg/models"
	"github.com/KapilSareen/go-project/pkg/types"
	"log"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	models.Connect()
	book := types.Book{
		Name:      "Book-for-Testing",
		Author:    "Unit test",
		Price:     69.69,
		OwnedFrom: time.Now().String(),
		IsAvail:   true,
	}

	err := models.AddBook(book)
	if err != nil {
		t.Errorf("failed to add book: %v", err)
		return
	}

	log.Println("Successfully added test book:", book.Name)

	models.Connect()
	err = models.DeleteBook(book.Name)
	if err != nil {
		t.Errorf("failed to delete book: %v", err)
		return
	}
	log.Println("Successfully deleted test book:", book.Name)
	fmt.Println("Test for adding book passed successfully")
}
