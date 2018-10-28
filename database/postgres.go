package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var db *sql.DB
var once sync.Once

func GetDataBase() *sql.DB {

	once.Do(func() {

		dbinf := fmt.Sprintf("user=%s password=%s dbname=%s host=127.0.0.1 port=5432 sslmode=disable", "kexibq", "22121996", "forumdb")
		var err error
		db, err = sql.Open("postgres", dbinf)
		if err != nil {
			log.Println("Can't connect to database", err)
		}
		err = db.Ping()
		if err != nil {
			log.Println("error in ping", err)
		}
	})
	return db
}

func CloseDB() {
	db.Close()
}
