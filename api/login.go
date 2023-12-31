package main

import (
	"encoding/json"
	"net/http"
	"sync"


	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)



var (
	// セッション情報を保存するためのマップ
	sessions = make(map[string]bool)
	// セッション情報へのアクセスを同期するためのミューテックス
	sessionMutex = &sync.Mutex{}
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// セッションIDを生成
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
	w.Write([]byte("login successfully!"))
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
	sessionMutex.Unlock()
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

	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}

	errMsg := ""

	// 空欄チェック
	if requestBody.Email == "" || requestBody.Password == "" {
		errMsg += "email、passwordは必須です。\n"
	}


}
