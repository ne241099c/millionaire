package ws

import (
	"encoding/json"
	"log"
	"millionaire/internal/game"
)

type GameAction struct {
	Client  *Client
	Message game.Message
}

type Hub struct {
	Game *game.Game

	Actions chan GameAction

	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),

		Game:    game.NewGame(),
		Actions: make(chan GameAction),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

			playerID := client.Conn.RemoteAddr().String()
			h.Game.Join(playerID)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}

		case action := <-h.Actions:
			h.handleGameMessage(action)
		}
	}
}

func (h *Hub) handleGameMessage(action GameAction) {
	playerID := action.Client.Conn.RemoteAddr().String()

	switch action.Message.Type {
	case game.MsgStartGame:
		log.Println("ゲーム開始リクエストを受信しました")
		h.Game.Start()      // カードを配る
		h.broadcastStatus() // 全員に知らせる

	case game.MsgPlayCard:
		var payload game.PlayCardPayload
		json.Unmarshal(action.Message.Payload, &payload)

		log.Printf("ゲーム処理: %s さんがカードを出そうとしています", playerID)

		if err := h.Game.PlayCard(playerID, payload.Cards); err != nil {
			log.Printf("❌ エラー: %v", err)
		} else {
			h.broadcastStatus()
		}

	case game.MsgJoin:
	}
}

func (h *Hub) broadcastStatus() {
	// h.Game.DebugPrint()

	for client := range h.Clients {
		playerID := client.Conn.RemoteAddr().String()

		var myHand []game.Card
		for _, p := range h.Game.Players {
			if p.ID == playerID {
				myHand = p.Hand
				break
			}
		}

		status := game.GameStatusPayload{
			Hand:        myHand,              // 手札
			TableCards:  h.Game.TableCards,   // 場のカード
			PlayerCount: len(h.Game.Players), // 参加人数
		}

		// JASONに変換
		payloadBytes, _ := json.Marshal(status)

		// メッセージ作成
		msg := game.Message{
			Type:    game.MsgGameStatus,
			Payload: payloadBytes, // RawMessage型に自動変換
		}

		// JASONにして送信
		msgBytes, _ := json.Marshal(msg)

		// 送信
		select {
		case client.Send <- msgBytes:
		default:
			close(client.Send)
			delete(h.Clients, client)
		}
	}

}
