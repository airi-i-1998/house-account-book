package conf

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func DBCon() *sql.DB {
	// DB接続
	db, err := sql.Open("mysql", "root:ajs2b0ti@tcp(localhost:3306)/house_account_book")
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	return DB
}
