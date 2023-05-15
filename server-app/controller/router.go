package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	// クロスオリジンリソース共有を有効化
	router.Use(cors.Default()) 

	//静的ファイル
	router.LoadHTMLGlob("../../view/*.html")

	//ホーム画面
	router.GET("/", func(c *gin.Context) {getTop(c, c.Writer)})
	//ログイン画面
	router.GET("/login", func(c *gin.Context) {getLogin(c, c.Writer)})
	router.POST("/login", func(c *gin.Context) {postLogin(c, c.Writer)})
	//サインアップ画面
	router.GET("/signup", func(c *gin.Context) {getSignup(c, c.Writer)})
	router.POST("/signup", func(c *gin.Context) {postSignup(c, c.Writer)})
	
	//LSプロトコル使用HTTPサーバー起動
	router.RunTLS(":8080", "../../certificate/server.pem", "../../certificate/server.key")

	return router
}