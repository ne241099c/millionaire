package main

import (
	"log"
	"net/http"

	"millionaire/internal/ws"
)

func main() {
	// 1. Hubを作成
	hub := ws.NewHub()

	// 2. Hubゴルーチンとして稼働させる
	go hub.Run()

	// 3. WebSocketの受付
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// 4. ファイル配布
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	log.Println("サーバー起動: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("起動失敗: ", err)
	}
}

// serveWs は、新しい接続があるたびに呼ばれる関数
// クライアントを作成し、Hubに登録
func serveWs(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	// wsパッケージで定義した Upgrader を使って接続をアップグレード
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 新しいClientを作成
	client := &ws.Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	// これで Hub.Run() の中の "case client := <-h.Register:" が動く
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
