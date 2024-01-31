package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()                // Ginルーターの新しいインスタンスを作成。HTTPリクエストを処理するためのルートやミドルウェアの設定を管理する。
	router.LoadHTMLGlob("template/*.html") // HTMLを読み込むディレクトリを指定

	data := "Hello Go/Gin!!"
	// ルートURL（/）に対するGETリクエストを処理するためのハンドラーを設定。
	// この関数は、リクエストが来た際に呼び出される。ここでは、ステータスコード200とともに index.html テンプレートをレンダリングする。
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data}) // map型で値を渡しているよ👀
	})
	// サーバーを起動し、デフォルトのポート（8080）で待ち受けを開始する
	router.Run()
}
