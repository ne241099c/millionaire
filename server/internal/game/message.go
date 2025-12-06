package game

import "encoding/json"

// MessageType は、通信の種類を表すラベル
// 文字列にしておくと、JSONを見た時にわかりやすい
type MessageType string

const (
	// クライアント -> サーバー
	MsgPlayCard MessageType = "play_card" // カードを出す
	MsgPass     MessageType = "pass"      // パスする
	MsgJoin     MessageType = "join"      // 参加したい

	MsgStartGame MessageType = "start_game" // ゲーム開始要求

	// サーバー -> クライアント
	MsgGameStatus MessageType = "game_status" // 現在の場の状況
	MsgError      MessageType = "error"       // エラー発生
)

type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type PlayCardPayload struct {
	Cards []Card `json:"cards"`
}

type GameStatusPayload struct {
	Hand        []Card `json:"hand"`         // あなたの手札
	TableCards  []Card `json:"table_cards"`  // 場に出ているカード
	IsMyTurn    bool   `json:"is_my_turn"`   // あなたの番？
	PlayerCount int    `json:"player_count"` // 参加人数
}
