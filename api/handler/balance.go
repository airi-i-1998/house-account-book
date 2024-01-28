package handler

import (
	"net/http"
	"example.com/m/api/conf"
	"time"
	"encoding/json"
	"log"
)

type RequestBalance struct{
	Description string `json:"description"`
	Date time.Time `json:"date"`
	Amount float64 `json:"amount"`
	Category string `json:"category"`
	Memo string `json:"memo"`
}

func Balance(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
			// CORSを許可する
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/Balance")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/Balance")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// リクエストボディをパース
	var requestBalance RequestBalance
	if err := json.NewDecoder(r.Body).Decode(&requestBalance); err != nil {
		log.Println("JSONデコードエラー:", err)
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	log.Println(requestBalance);

	// データベースに挿入
	stmt, err := conf.DB.Prepare("INSERT INTO transactions (description, date, amount, category, memo) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
			http.Error(w, "データベースへの保存に失敗しました", http.StatusInternalServerError)
			return
	}
	defer stmt.Close()

	_, err = stmt.Exec(requestBalance.Description, requestBalance.Date, requestBalance.Amount, requestBalance.Category, requestBalance.Memo)
	if err != nil {
			http.Error(w, "データベースへの保存に失敗しました", http.StatusInternalServerError)
			return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusCreated)
}
