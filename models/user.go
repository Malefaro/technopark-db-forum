package models

import (
	"database/sql"
	"log"
)

type User struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname"`
}

func CreateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("INSERT INTO users (about, email, fullname, nickname) VALUES ($1, $2, $3, $4)", user.About, user.Email, user.Fullname, user.Nickname)
	if err != nil {
		log.Printf("Function: CreateUser, User: %v, Error: %v",user, err)
		return err
	}
	return nil
}