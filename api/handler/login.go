package handler

import (
	"encoding/json"
	"database/sql"
	"net/http"
	"sync"

	"example.com/m/api/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// RequestLogin はPOSTリクエストのJSONデータを格納する構造体。
type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	// セッション情報を保存するためのマップ
	sessions = make(map[string]bool)
	// セッション情報へのアクセスを同期するためのミューテックス
	sessionMutex = &sync.Mutex{}
)

// データベースからユーザー情報を取得する
func getUser(email, password string) (string, error) {
	var storedPassword string
	err := conf.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	//エラー発生確認
	if err != nil {
			//エラー発生の場合
			if err == sql.ErrNoRows {
					return "", nil // ユーザーが見つからない場合
			}
			//エラーがない場合
			return "", err // その他のエラーが発生した場合
	}
	return storedPassword, nil // パスワードを返す
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin RequestLogin
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
			http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
			return
	}

	email := requestLogin.Email
	password := requestLogin.Password

	// データベース内のユーザー情報を取得
	storedPassword, err := getUser(email, password)
	if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	}

	// パスワードが一致するかを確認
	if password != storedPassword {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
	}

	sessionID := uuid.New().String()

	//サーバサイドでsessionを保存
	sessionMutex.Lock()
	sessions[sessionID] = true
	sessionMutex.Unlock()

	// クライアントにセッションを送信
	http.SetCookie(w, &http.Cookie{
			Name:   "session_id",
			Value:  sessionID,
			Path:   "/",
			MaxAge: 7200, // 有効期限を2時間に設定
	})

	w.Write([]byte("Login successful!"))
}
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// CookieからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionMutex.Lock()
	authenticated := sessions[cookie.Value]
	sessionMutex.Unlock()

	// セッションが有効かチェック
	if !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Write([]byte("Welcome to your dashboard!"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// CookieからセッションIDを取得し、サーバーサイドでセッションを削除
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie.Value)
		sessionMutex.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Cookieを削除
	})

	w.Write([]byte("Logout successful!"))
}
