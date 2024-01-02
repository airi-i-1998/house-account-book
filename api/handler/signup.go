package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"example.com/m/api/conf"
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

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := ErrorResponse{
		Error: message,
	}

	// エラーメッセージをJSON形式にシリアライズ
	responseJSON, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, "JSONのシリアライズに失敗しました", http.StatusInternalServerError)
		return
	}

	// Content-Typeを設定してHTTPレスポンスのボディにJSON形式のエラーメッセージを設定
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "レスポンスの書き込みに失敗しました", http.StatusInternalServerError)
		return
	}
}

func Singup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		// CORSを許可する
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// リクエストボディの読み取り
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
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
		fmt.Println(errMsg)
		sendErrorResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	// 受け取ったデータを使って何らかの処理を行う（ここでは簡単な例としてそのままレスポンスを作成）
	responseMessage := fmt.Sprintf("Hello, %s! Your email is %s. Password is %s", requestBody.Name, requestBody.Email, requestBody.Password)

	stmt, err := conf.DB.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
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
