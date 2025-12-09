package ws

import (
	"encoding/json"
	"log"
	"millionaire/internal/game"
)

type Room struct {
	ID   string
	Game *game.Game

	Actions chan GameAction

	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewRoom(id string) *Room {
	return &Room{
		ID:         id,
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Game:       game.NewGame(),
		Actions:    make(chan GameAction),
	}
}

func (r *Room) Run() {
	log.Printf("🏠 部屋 [%s] が起動しました", r.ID)
	for {
		select {
		case client := <-r.Register:
			r.Clients[client] = true

			playerID := client.Conn.RemoteAddr().String()
			r.Game.Join(playerID, client.Name)

			r.broadcastStatus()

		case client := <-r.Unregister:
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.Send)
			}

		case message := <-r.Broadcast:
			for client := range r.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Clients, client)
				}
			}

		case action := <-r.Actions:
			r.handleGameMessage(action)
		}
	}
}

func (r *Room) handleGameMessage(action GameAction) {
	playerID := action.Client.Conn.RemoteAddr().String()

	switch action.Message.Type {
	case game.MsgStartGame:
		if r.Game.IsActive {
			log.Println("ゲーム中なので開始リクエストを無視しました")
			break
		}
		log.Println("ゲーム開始リクエストを受信しました")
		r.Game.Start()
		r.broadcastStatus()

	case game.MsgPlayCard:
		var payload game.PlayCardPayload
		json.Unmarshal(action.Message.Payload, &payload)

		log.Printf("ゲーム処理: %s さんがカードを出そうとしています", playerID)

		if err := r.Game.PlayCard(playerID, payload.Cards); err != nil {
			log.Printf("❌ エラー: %v", err)
		} else {
			r.broadcastStatus()
		}

	case game.MsgJoin:

	case game.MsgPass:
		log.Printf("ゲーム処理: %s さんがパスしました", playerID)

		if err := r.Game.Pass(playerID); err != nil {
			log.Printf("❌ エラー: %v", err)
		} else {
			r.broadcastStatus()
		}
	}
}

func (r *Room) broadcastStatus() {
	var allPlayersData []game.PlayerData
	for _, p := range r.Game.Players {
		allPlayersData = append(allPlayersData, game.PlayerData{
			ID:   p.ID,
			Hand: p.Hand,
			Rank: p.Rank,
		})
	}

	for client := range r.Clients {
		playerID := client.Conn.RemoteAddr().String()

		var myHand []game.Card
		var amIActivePlayer bool

		for _, p := range r.Game.Players {
			if p.ID == playerID {
				myHand = p.Hand
				if len(p.Hand) > 0 {
					amIActivePlayer = true
				}
				break
			}
		}

		isSpectator := r.Game.IsActive && !amIActivePlayer

		currentPlayer := r.Game.Players[r.Game.TurnIndex]
		IsMyTurn := (currentPlayer.ID == playerID)
		effectiveRev := (r.Game.IsRevolution != r.Game.Is11Back)

		// ステータス作成
		status := game.GameStatusPayload{
			Hand:         myHand,              // 手札
			TableCards:   r.Game.TableCards,   // 場のカード
			PlayerCount:  len(r.Game.Players), // 参加人数
			IsMyTurn:     IsMyTurn,            // 自分の番?
			IsRevolution: effectiveRev,        // 革命中？
		}

		// 観戦者なら全員のデータを添付する
		if isSpectator {
			status.AllPlayers = allPlayersData
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
			delete(r.Clients, client)
		}
	}

}
