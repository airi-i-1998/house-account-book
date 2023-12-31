package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ルーティング設定
	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/logout", logoutHandler)

	// サーバーを起動し、ポート8080でリクエストを待ち受ける
	port := ":8080"
	fmt.Printf("サーバーをポート%sで起動中...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %s", err)
	}
}
