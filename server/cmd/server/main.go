package main

import (
	"fmt"
	"millionaire/internal/game"
	"net/http"
)

func main() {
	fmt.Println("--- 階段判定テスト ---")

	c3 := game.NewCard(game.Spade, game.Three)
	c4 := game.NewCard(game.Spade, game.Four)
	c5 := game.NewCard(game.Spade, game.Five)
	c7 := game.NewCard(game.Spade, game.Seven) // 6が抜けている
	joker := game.NewCard(game.Joker, 0)

	// ケース1: バラバラの順番 (5, 3, 4) -> ソートして判定できるか？
	seq1 := []game.Card{c5, c3, c4}
	fmt.Printf("1. [♠5, ♠3, ♠4] は階段？ -> %v\n", game.IsSequence(seq1))

	// ケース2: マーク違い (♠3, ♥4, ♠5) -> Falseになるか？
	c4_heart := game.NewCard(game.Heart, game.Four)
	seq2 := []game.Card{c3, c4_heart, c5}
	fmt.Printf("2. [♠3, ♥4, ♠5] は階段？ -> %v\n", game.IsSequence(seq2))

	// ケース3: 穴あき階段とジョーカー (3, Joker, 5) -> Jokerが4の代わりになるか？
	// 差が2あるので、Jokerを1枚消費して成立するはず
	seq3 := []game.Card{c3, joker, c5}
	fmt.Printf("3. [♠3, Joker, ♠5] は階段？ -> %v\n", game.IsSequence(seq3))

	// ケース4: 穴が大きすぎる (3, Joker, 7) -> Joker1枚じゃ埋まらない (4,5,6が必要)
	seq4 := []game.Card{c3, joker, c7}
	fmt.Printf("4. [♠3, Joker, ♠7] は階段？ -> %v\n", game.IsSequence(seq4))

	fmt.Println("--------------------")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server OK")
	})
	fmt.Println("Server Start...")
	http.ListenAndServe(":8080", nil)
}
