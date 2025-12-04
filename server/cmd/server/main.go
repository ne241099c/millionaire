package main

import (
	"fmt"
	"net/http"

	// 作った game パッケージを読み込みます
	"millionaire/internal/game"
)

func main() {
	// 1. 山札を作成
	fmt.Println("1. 山札を新品で作成します...")
	deck := game.NewDeck(1)
	fmt.Printf("   -> 合計枚数: %d枚\n", len(deck))

	// 2. シャッフル
	fmt.Println("2. シャッフルします...")
	deck.Shuffle()

	// 3. 5枚引いてみる
	fmt.Println("3. 上から5枚引いてみます...")
	hand := deck.Draw(5)

	// 4. 引いたカードを画面に出力
	fmt.Println("--- 引いたカード ---")
	for _, c := range hand {
		fmt.Printf("%s ", c)
	}
	fmt.Println("\n-------------------")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "大富豪サーバー、稼働中！")
	})

	fmt.Println("サーバーがポート8080で起動しました")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("起動エラー:", err)
	}
}
