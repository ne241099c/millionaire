package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/" にアクセスが来たら処理する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "大富豪サーバー、稼働中！")
	})

	fmt.Println("サーバーがポート8080で起動しました...")
	// ここでサーバーが待機状態になり、プログラムが終了しなくなります
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("起動エラー:", err)
	}
}
