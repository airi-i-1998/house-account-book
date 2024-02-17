package handler

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
)

var (
	// セッション情報を保存するためのマップ(検索と取得)
	sessions = make(map[string]bool)
	// セッション情報へのアクセスを同期するためのミューテックス
	sessionMutex = &sync.Mutex{}
)

// generateRandomString は指定された長さのランダムな文字列を生成します。
func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// 本来はIDとパスワードでのユーザー認証のロジックを実装する

	// セッションIDを生成(32文字のランダムな文字列を生成)
	sessionID,err := generateRandomString(32)
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error generating session ID:", err)
		return
	}

	// サーバーサイドでセッションを保存
	sessionMutex.Lock()
	sessions[sessionID] = true
	for id := range sessions {
		fmt.Println("loginHandler: Current session ID is ", id)
	}
	sessionMutex.Unlock()

	// クライアントにセッションIDをCookieとして送信
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  sessionID,
		Path:   "/",
		MaxAge: 60, // 有効期限を60秒（1分）に設定
	})

	w.Write([]byte("login successfully!"))

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// CookieからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionMutex.Lock()
	authenticated, ok := sessions[cookie.Value]
	if ok {
		for id := range sessions {
			fmt.Println("loginHandler: Current session ID is ", id)
		}
	}
	sessionMutex.Unlock()

	// セッションが有効かチェック
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	w.Write([]byte("login successfully!"))
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	// CookieからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	sessionMutex.Lock()
	authenticated, ok := sessions[cookie.Value]
	if ok {
		for id := range sessions {
			fmt.Println("balanceHandler: Current session ID is ", id)
		}
	}
	sessionMutex.Unlock()

	// セッションが有効かチェック
	if !ok || !authenticated {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	w.Write([]byte("Welcome to your balance!"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// CookieからセッションIDを取得し、サーバーサイドでセッションを削除
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionMutex.Lock()
		delete(sessions, cookie.Value)
		if len(sessions) == 0 {
			fmt.Println("logoutHandler: Current session ID is nothing")
		}
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
