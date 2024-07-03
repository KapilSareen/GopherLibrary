package models

import (
	"context"
	"fmt"
	"log"
)

func Issue(book string, username string) {
	query := "UPDATE `book` SET issued_to=?, is_avail=false WHERE name=?"
	insertResult, err := db.ExecContext(context.Background(), query, username, book)
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		fmt.Print(err.Error())
	}

	db.Close()
}

func BookAvail(name string) bool {
	var bookname string
	if err := db.QueryRow("SELECT name from book where name = ? and is_avail = 1",
		name).Scan(&bookname); err != nil {
		return false
	}
	db.Close()
	return bookname != ""

}

func BookExist(name string) bool {
	var bookname string
	if err := db.QueryRow("SELECT name from book where name = ?",
		name).Scan(&bookname); err != nil {
		return false
	}
	db.Close()
	return bookname != ""

}

func DeleteBook(book string)  {
	query := "DELETE FROM book WHERE name=?"
	insertResult, err := db.ExecContext(context.Background(), query, book)
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		log.Fatal(err) 
	}
	db.Close()
}
