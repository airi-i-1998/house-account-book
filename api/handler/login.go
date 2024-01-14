package handler

import (
	"log"
	"encoding/json"
	"net/http"
	"sync"
	"database/sql"

	"example.com/m/api/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
func getUser(email string) (string, error) {
	var storedPassword string
	err := conf.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // ユーザーが見つからない場合
		}
		return "", err // その他のエラーが発生した場合
	}
	return storedPassword, nil // ハッシュ化されたパスワードを返す
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// CORSを許可する
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var requestLogin RequestLogin
	err := json.NewDecoder(r.Body).Decode(&requestLogin)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}

	email := requestLogin.Email
	password := requestLogin.Password

	// データベース内のユーザー情報を取得
	storedPassword, err := getUser(email)
	if err != nil {
		log.Printf("Login failed for email: %s, error: %s",storedPassword, err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Retrieved stored password for email: %s", storedPassword)

	// データベースから取得したハッシュ化されたパスワードと、入力された生のパスワードを比較
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Printf("Login failed for email: %s, error: %s",password, err.Error())
		return
	}

	sessionID := uuid.New().String()

	// サーバサイドでsessionを保存
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
	// セッションIDをCookieから取得
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
	// セッションIDをCookieから取得し、サーバーサイドでセッションを削除
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
