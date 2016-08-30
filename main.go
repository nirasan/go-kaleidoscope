package main

import (
	"github.com/nirasan/go-kaleidoscope/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	// 万華鏡画像表示 | クエリパラメータの w で画像サイズ(ピクセル) n で生成要素数(多いほどごちゃっとする)の指定
	http.HandleFunc("/", handler.Kaleidoscope)
	// 砂嵐画像表示
	http.HandleFunc("/mono", handler.Mono)
	// 縞画像表示
	http.HandleFunc("/stripe", handler.Stripe)
	// 縞画像の回転
	http.HandleFunc("/rotate", handler.Rotate)
	// 縞画像の扇形に切り抜いて回転してコピー
	http.HandleFunc("/sector", handler.Sector)
	// sector でピクセルの欠けが出たので修正版
	http.HandleFunc("/sector2", handler.Sector2)
	// 環境変数から公開ポートの指定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// サーバー起動
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
