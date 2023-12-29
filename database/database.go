package database

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID int `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

func database() {
	// dsn := os.Getenv("DSN")
	//DBに接続する
	db, err := sqlx.Open("mysql","root:ajs2b0ti@tcp(localhost:3306)/house_account_book")
	if err != nil {
		//現在の関数が終了する時点で、データベースの接続をクローズすること
		// defer db.Close()
    log.Fatal(err)
	}

	//SQL実行
	rows,err := db.Queryx("SELECT id,name,email,password FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	results := make([]User,0)
	//データを一行ずつ取得
	for rows.Next(){
		var user User


//user変数にクエリ結果を格納
		err := rows.StructScan(&user)

		if err != nil{
			log.Fatal(err)
		}

		results = append(results, user)
	}

	fmt.Println(results)

	// テーブルのマイグレーション
	// db.AutoMigrate(&User{})

}
