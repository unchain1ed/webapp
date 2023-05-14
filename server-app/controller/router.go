package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	
	router.LoadHTMLGlob("../../view/*.html")

	// // 静的ファイルのパスを指定
	// router.Static("/view", "../view/home.html")

	// // ルーターの設定
	// // URLへのアクセスに対して静的ページを返す
	// router.StaticFS("/unchain1ed", http.Dir("../view"))

	// router.GET("/", getTop)
	// router.GET("/signup", getSignup)
	// router.POST("/signup", postSignup)
	// router.GET("/login", getLogin)
	// router.POST("/login", postLogin)

	// // 静的ファイルを提供する
	// router.Static("/", "../../frontend-app/public")

	// // SPAのために、すべてのルートにindex.htmlを返すように設定する
	// router.NoRoute(func(c *gin.Context) {
	// 	c.File("../../frontend-app/public/home.html")
	// })

	// router.LoadHTMLGlob("../../view/*.html")
	router.GET("/", func(c *gin.Context) {
		w := c.Writer
		getTop(c, w)
	})
	// router.GET("/", getTop) // ログイン画面（GET処理）
	router.RunTLS(":8080", "../../certificate/server.pem", "../../certificate/server.key")

	return router
}