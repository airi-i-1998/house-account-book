package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

// RequestBody はPOSTリクエストのJSONデータを格納する構造体。
type RequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ResponseBody はレスポンスのJSONデータを格納する構造体。
type ResponseBody struct {
	Message string `json:"message"`
}

func main() {
	// ルーティング設定
	http.HandleFunc("/signup", handleAPI)

	// サーバーを起動し、ポート8080でリクエストを待ち受ける
	port := ":8080"
	fmt.Printf("サーバーをポート%sで起動中...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %s", err)
	}
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	// POSTメソッド以外のリクエストを受け付けないようにする
	if r.Method != http.MethodPost {
		http.Error(w, "POSTメソッドのみ受け付けています", http.StatusMethodNotAllowed)
		return
	}

	//DB接続
	db, err := sql.Open("mysql", "root:ajs2b0ti@tcp(localhost:3306)/house_account_book")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// リクエストボディの読み取り
	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}

	errMsg := ""

// 空欄チェック
if requestBody.Name == "" || requestBody.Email == "" || requestBody.Password == "" {
    errMsg += "name、email、passwordは必須です。\n"
}

// emailフォーマットチェック
if requestBody.Email != "" && !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(requestBody.Email) {
    errMsg += "有効なメールアドレスを入力してください。\n"
}

// パスワードの文字数制限と正規表現を使用したパスワードのバリデーション
if requestBody.Password != "" && (len(requestBody.Password) < 6 || len(requestBody.Password) > 20 || !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(requestBody.Password)) {
    errMsg += "passwordは6文字以上20文字以下の英数字で入力してください。\n"
}

// エラーメッセージがある場合、HTTPレスポンスで返す
if errMsg != "" {
    http.Error(w, errMsg, http.StatusBadRequest)
    return
}

	// 受け取ったデータを使って何らかの処理を行う（ここでは簡単な例としてそのままレスポンスを作成）
	responseMessage := fmt.Sprintf("Hello, %s! Your email is %s. Password is %s", requestBody.Name, requestBody.Email, requestBody.Password)

	stmt, err := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "データベースへの保存に失敗しました", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(requestBody.Name, requestBody.Email, requestBody.Password)
	if err != nil {
		http.Error(w, "データベースへの保存に失敗しました", http.StatusInternalServerError)
		return
	}

	// レスポンスの作成
	responseBody := ResponseBody{
		Message: responseMessage,
	}

	// レスポンスをJSON形式で返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "レスポンスの作成に失敗しました", http.StatusInternalServerError)
		return
	}
}
