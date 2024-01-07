package handler

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/api/conf"
	"example.com/m/api/handler"
)

func main() {
	db := conf.DBCon()
	defer db.Close()

	// ルーティング設定
	http.HandleFunc("/signup", handler.Signup)
	http.HandleFunc("/login", handler.Login)
	// http.HandleFunc("/dashboard", dashboardHandler)
	// http.HandleFunc("/logout", logoutHandler)

	// サーバーを起動し、ポート8080でリクエストを待ち受ける
	port := ":8080"
	fmt.Printf("サーバーをポート%sで起動中...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %s", err)
	}
}
