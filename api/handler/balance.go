package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"example.com/m/api/conf"
)

type RequestBalance struct {
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Memo        string    `json:"memo"`
}

func Balance(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// CORSを許可する
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// リクエストボディをパース
	var requestBalance RequestBalance
	if err := json.NewDecoder(r.Body).Decode(&requestBalance); err != nil {
		log.Println("JSONデコードエラー:", err)
		http.Error(w, "リクエストの解析に失敗しました", http.StatusBadRequest)
		return
	}
	log.Println(requestBalance)

	dateStr := r.FormValue("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println("日付のパースエラー:", err)
		http.Error(w, "日付のパースに失敗しました", http.StatusBadRequest)
		return
	}

	// パースした日付をUTCに設定
	requestBalance.Date = date.UTC()

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
