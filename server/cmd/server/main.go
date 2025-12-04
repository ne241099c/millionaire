package main

import (
	"fmt"
	"millionaire/internal/game"
	"net/http"
)

func main() {
	c3 := game.NewCard(game.Spade, game.Three)
	c2 := game.NewCard(game.Club, game.Two)
	joker := game.NewCard(game.Joker, 0)

	fmt.Println("--- 革命テスト (Revolution!) ---")

	// 革命状態をONにする
	isRevolution := true

	// テスト1: 3 vs 2
	// 3の強さ: -1, 2の強さ: -14
	if game.IsStronger(c3, c2, isRevolution) {
		fmt.Printf("革命中: %s は %s より強い -> OK!\n", c3, c2)
	} else {
		fmt.Printf("革命中: %s は %s より弱い -> NG...\n", c3, c2)
	}

	// テスト2: Joker vs 3
	if game.IsStronger(joker, c3, isRevolution) {
		fmt.Printf("革命中: %s は %s より強い -> OK!\n", joker, c3)
	} else {
		fmt.Printf("革命中: %s は %s より弱い -> NG...\n", joker, c3)
	}

	fmt.Println("------------------------------")

	http.ListenAndServe(":8080", nil)
}
