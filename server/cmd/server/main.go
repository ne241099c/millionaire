package main

import (
	"log"
	"net/http"

	"millionaire/internal/ws"
)

func main() {
	// 1. Lobbyを作成
	lobby := ws.NewLobby()

	// 3. WebSocketの受付
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// ロビーに案内を任せる
		lobby.ServeWs(w, r)
	})

	// 4. ファイル配布
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("サーバー起動: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("起動失敗: ", err)
	}
}
