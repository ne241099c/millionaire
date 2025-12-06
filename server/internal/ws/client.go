package ws

import (
	"encoding/json"
	"log"
	"millionaire/internal/game"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// 1. メッセージを受け取る
		_, messageData, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		// 2. TypeとPayloadを解析する
		var msg game.Message
		if err := json.Unmarshal(messageData, &msg); err != nil {
			log.Printf("JSON解析エラー: %v", err)
			continue // 変なデータが来たら無視して次へ
		}

		// 3. Typeによって処理を分ける
		switch msg.Type {

		case game.MsgPlayCard:
			// "play_card" の場合、Payloadの中身は PlayCardPayload のはず
			var payload game.PlayCardPayload

			// Payloadを、PlayCardPayload型として解析し直す
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				log.Printf("ペイロード解析エラー: %v", err)
				break
			}

			action := GameAction{
				Client:  c,
				Message: msg,
			}
			c.Hub.Actions <- action

		case game.MsgStartGame:
			action := GameAction{
				Client:  c,
				Message: msg,
			}
			c.Hub.Actions <- action

		case game.MsgPass:
			log.Println("★パスされました")

		default:
			log.Printf("知らないメッセージタイプです: %s", msg.Type)
		}

		// 動作確認のために全員にそのままオウム返ししておく
		c.Hub.Broadcast <- messageData
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for {
		message, ok := <-c.Send
		if !ok {
			return
		}

		writer, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		writer.Write(message)

		if err := writer.Close(); err != nil {
			return
		}
	}
}
