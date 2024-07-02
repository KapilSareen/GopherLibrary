package models

import (
	"context"
	"fmt"
	"github.com/KapilSareen/go-project/pkg/types"
	"log"
)

func Signup(name string, password string) error {
	query := "INSERT INTO `users` ( `name`, `password`) VALUES ( ?,?)"
	insertResult, err := db.ExecContext(context.Background(), query, name, password)
	if err != nil {
		return err
	}
	id, _ := insertResult.LastInsertId()

	log.Printf("inserted id: %d", id)
	db.Close()
	return nil

}

func Login(name string, password string) (types.User, error) {
	var user types.User
	if err := db.QueryRow("SELECT * from users where name = ? and password = ?",
		name, password).Scan(&user.ID, &user.Name, &user.IsAdmin,
		&user.Password); err != nil {
		return user, err
	}
	db.Close()
	return user, nil
}

func IsAdmin(user types.User) bool {
	var IsAdmin bool
	if err := db.QueryRow("SELECT IsAdmin from users where name = ? and password = ?",
		user.Name, user.Password).Scan(&IsAdmin); err != nil {
		fmt.Print(err)
		return IsAdmin
	}
	db.Close()
	return IsAdmin
}
