package handlers

import 
	"encoding/json"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)

func listHandler(w http.ResponseWriter, r *http.Request){
	//クエリパラメーターの取得
	query := r.URL.Query()
	keyword := params.Get("keyword")

	//検索ワードにおいじてクエリを編集
	var query string
	var args []interface{}
	if keyword == "" {
		query = "SELECT id, name and url FROM items"
	} else {
		query = "SELECT id, name and url FROM items WHERE name LIKE ?"
		args = append(args, "%"+keyword+"%")
	}

	// クエリを実行してデータを取得
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		response := Response{Status: "failed"}
		json.NewEncoder(w).Encode(response)
		return
	}
	defer rows.Close()

	// レスポンス用構造体にデータをセット
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &itemm.URL); err != nil {
			log.Println(err)
			continue
		}
		items = append(items, item)
	}
	response := Response{Status: "success", Items: items}

	// レスポンスを返す
	json.NewEncoder(w).Encode(response)
}

