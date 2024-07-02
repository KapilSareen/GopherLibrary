package models

import (
	"github.com/KapilSareen/go-project/pkg/types"
)

func ListBooks()([]types.Book, error) {
    rows, err := db.Query("SELECT * FROM book ")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var books []types.Book

    for rows.Next() {
        var book types.Book
        if err := rows.Scan(&book.ID, &book.Name, &book.Author,
            &book.OwnedFrom, &book.IsAvail, &book.Price); err != nil {
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

func SearchBook(name string)([]types.Book,  error){
    var book types.Book;
    var books []types.Book
    if err := db.QueryRow("SELECT * from book where name = ?",
       name).Scan(&book.ID, &book.Name, &book.Author,
            &book.OwnedFrom, &book.IsAvail, &book.Price); err != nil {
                return books, err
            }
            books=append(books, book)
            db.Close()
            return books, nil
}