package models

import (
	"github.com/KapilSareen/go-project/pkg/types"
)

func ListBooks() ([]types.Book, error) {
	rows, err := db.Query("SELECT * FROM book ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book

	for rows.Next() {
		var book types.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author,
			&book.OwnedFrom, &book.IsAvail, &book.Price, &book.Issued_to ); err != nil {
			return books, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return books, err
	}
	db.Close()
	return books, nil
}

func SearchBook(name string) ([]types.Book, error) {
	var books []types.Book
	rows, err := db.Query("SELECT * FROM book WHERE name LIKE ?", name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book types.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.OwnedFrom, &book.IsAvail, &book.Price, &book.Issued_to); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}