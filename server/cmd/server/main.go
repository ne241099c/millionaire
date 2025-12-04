package main

import (
	"fmt"
	"millionaire/internal/game"
	"net/http"
)

func main() {
	fmt.Println("--- ペア判定テスト ---")

	// テスト用カード
	c3_spade := game.NewCard(game.Spade, game.Three)
	c3_heart := game.NewCard(game.Heart, game.Three)
	c4_dia := game.NewCard(game.Diamond, game.Four)
	joker := game.NewCard(game.Joker, 0)

	// ケース1: 3のペア (3, 3) -> 成功するはず
	pair1 := []game.Card{c3_spade, c3_heart}
	fmt.Printf("1. [♠3, ♥3] はペア？ -> %v\n", game.IsPair(pair1))

	// ケース2: 3とJoker (3, Joker) -> 成功するはず
	pair2 := []game.Card{c3_spade, joker}
	fmt.Printf("2. [♠3, Joker] はペア？ -> %v\n", game.IsPair(pair2))

	// ケース3: バラバラ (3, 4) -> 失敗するはず
	pair3 := []game.Card{c3_spade, c4_dia}
	fmt.Printf("3. [♠3, ♦4] はペア？ -> %v\n", game.IsPair(pair3))

	// ケース4: 3枚ペア (3, 3, Joker) -> 成功するはず
	pair4 := []game.Card{c3_spade, c3_heart, joker}
	fmt.Printf("4. [♠3, ♥3, Joker] はペア？ -> %v\n", game.IsPair(pair4))

	fmt.Println("--------------------")

	// サーバー起動処理
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server OK")
	})
	fmt.Println("Server Start...")
	http.ListenAndServe(":8080", nil)
}
