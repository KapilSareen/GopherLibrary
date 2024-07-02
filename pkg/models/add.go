package models

import (
	"context"
	"github.com/KapilSareen/go-project/pkg/types"
	"log"
)

func AddBook(book types.Book) {

	query := "INSERT INTO `book` ( `name`, `author`,`price`) VALUES ( ?, ?,?)"
	insertResult, err := db.ExecContext(context.Background(), query, book.Name, book.Author, book.Price)
	if err != nil {
		log.Fatalf("impossible insert : %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)

	db.Close()
}
