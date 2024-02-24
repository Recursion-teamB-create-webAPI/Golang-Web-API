package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // クエリパラメータを解析する
    query := r.URL.Query()
    name := query.Get("name") // "名前"から"name"へ修正

    // レスポンス用のマップを作成
    response := map[string]string{
        "message": "Hello " + name, // “message”： "Hello " + name、から修正
    }

    // Content-Typeヘッダーをapplication/jsonに設定
    w.Header().Set("Content-Type", "application/json")

    // マップをJSONにエンコードしてレスポンスとして送信
    json.NewEncoder(w).Encode(response)
}

func main() {
    fmt.Println("Starting the server!")
    
    // ルートとハンドラ関数を定義
    http.HandleFunc("/api/hello", helloHandler)

    // 8000番ポートでサーバを開始
    http.ListenAndServe(":8000", nil)
}