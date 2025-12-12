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

type PlayerData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Hand      []Card `json:"hand"`
	HandCount int    `json:"hand_count,omitempty"`
	Rank      int    `json:"rank"`
}

type GameStatusPayload struct {
	Hand          []Card       `json:"hand"`            // あなたの手札
	TableCards    []Card       `json:"table_cards"`     // 場に出ているカード
	IsMyTurn      bool         `json:"is_my_turn"`      // あなたの番？
	PlayerCount   int          `json:"player_count"`    // 参加人数
	IsRevolution  bool         `json:"is_revolution"`   // 革命中？
	AllPlayers    []PlayerData `json:"all_players"`     // 全プレイヤーの情報
	WinnerName    string       `json:"winner_name"`     // 勝者の名前
	IsActive      bool         `json:"is_active"`       // ゲームがアクティブかどうか
	CurrentTurnID string       `json:"current_turn_id"` // 現在のターンのプレイヤーID
}
