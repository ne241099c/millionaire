package ws

import (
	"log"
	"millionaire/internal/game"
	"net/http"
)

type GameAction struct {
	Client  *Client
	Message game.Message
}

type Lobby struct {
	Rooms map[string]*Room
}

func NewLobby() *Lobby {
	return &Lobby{
		Rooms: make(map[string]*Room),
	}
}

func (l *Lobby) CreateRoom(roomID string) *Room {
	if room, ok := l.Rooms[roomID]; ok {
		return room
	}

	// 新しい部屋を作成
	newRoom := NewRoom(roomID)
	l.Rooms[roomID] = newRoom

	go newRoom.Run()

	log.Printf("🏢 ロビー: 新しい部屋 [%s] を作成しました", roomID)
	return newRoom
}

func (l *Lobby) ServeWs(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから部屋IDを取得
	query := r.URL.Query()
	roomID := r.URL.Query().Get("room")
	playerName := query.Get("name")
	if roomID == "" {
		roomID = "default" // 指定がなければ "default" 部屋へ
	}
	if playerName == "" {
		playerName = "名無し" // ★デフォルト値
	}

	// 部屋を取得または作成
	room := l.CreateRoom(roomID)

	// WebSocket接続
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 新しいClientを作成
	client := &Client{
		Room: room,
		Conn: conn,
		Send: make(chan []byte, 256),
		Name: playerName,
	}

	// 部屋に入室させる
	client.Room.Register <- client

	// 読み書き開始
	go client.WritePump()
	go client.ReadPump()
}
