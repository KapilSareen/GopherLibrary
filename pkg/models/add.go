package models

import (
	"context"
	"github.com/KapilSareen/go-project/pkg/types"
)

func AddBook(book types.Book) error {

	query := "INSERT INTO `book` ( `name`, `author`,`price`) VALUES ( ?, ?,?)"
	insertResult, err := db.ExecContext(context.Background(), query, book.Name, book.Author, book.Price)
	if err != nil {
		return err
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
