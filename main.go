package main

import (
	"github.com/daisuzuki829/hello_route/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// テンプレートをロード
	r.LoadHTMLGlob("templates/*.html")

	// ハンドラの指定
	r.GET("/hello", routes.Hello)
	r.GET("/hello/:name", routes.HelloName)

	// グルーピング
	api := r.Group("/api")
	{
		api.GET("/hello", routes.HelloJson)
		api.GET("/hello/:name", routes.HelloJsonName)
	}

	// どのルーティングにも当てはまらなかった場合
	r.NoRoute(routes.NoRoute)
	r.Run(":8080")
}
