package controller

import (
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/unchain1ed/server-app/model/redis"
	"github.com/unchain1ed/server-app/controller/blog"
	"github.com/unchain1ed/server-app/controller/login"
	"github.com/unchain1ed/server-app/controller/logout"
	"github.com/unchain1ed/server-app/controller/setting"
	"github.com/unchain1ed/server-app/controller/common"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func init() {
	//環境変数設定
	//main.goからの相対パス指定
	envErr := godotenv.Load("./build/app/.env")
	if envErr != nil {
		fmt.Println("Error loading .env file", envErr)
	}
}	

//APIエンドポイントとクロスオリジンリソース共有（CORS）の設定
func GetRouter() *gin.Engine {
	//ルターを定義
	router := gin.Default()

	// クロスオリジンリソース共有_CORS設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://server-app:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"Cookie",
	}
	config.AllowCredentials = true
	//クロスオリジンリソース共有を有効化
	router.Use(cors.New(config))

	//***ホーム概要画面***
	router.GET("/", isAuthenticated(), func(c *gin.Context) {blog.GetTop(c)})

	//***ログイン画面***
	router.GET("/login", func(c *gin.Context) {login.GetLogin(c)})
	router.POST("/login", func(c *gin.Context) {login.PostLogin(c)})

	//***ブログ概要画面***
	router.POST("/blog/post", isAuthenticated(), func(c *gin.Context) {blog.PostBlog(c)})
	//BlogOverview画面
	router.GET("/blog/overview", isAuthenticated(),func(c *gin.Context) {blog.GetBlogOverview(c)})
	//BlogIDによるView画面
	router.GET("/blog/overview/post/:id", isAuthenticated(), func(c *gin.Context) {blog.GetBlogViewById(c)})
	//ブログ記事編集API
	router.POST("/blog/edit", isAuthenticated(), func(c *gin.Context) {blog.PostEditBlog(c)})
	//ブログ記事消去API
	router.GET("/blog/delete/:id", isAuthenticated(), func(c *gin.Context) {blog.GetDeleteBlog(c)})

	//***ID情報編集画面***
	//ID変更API
	router.POST("/update/id", isAuthenticated(), func(c *gin.Context) {setting.PostUpdateId(c)})

	//***PW情報編集画面***
	//PW変更API
	router.POST("/update/pw", isAuthenticated(), func(c *gin.Context) {setting.PostUpdatePw(c)})

	//***ログアウト画面***
	//ログアウト実行API
	router.POST("/logout", isAuthenticated(), func(c *gin.Context) {logout.DecideLogout(c)})

	//***会員情報登録画面***
	//登録画面遷移
	router.POST("/regist", isAuthenticated(), func(c *gin.Context) {blog.PostRegist(c)})

	//***共通API***
	//セッションからログインIDを取得するAPI
	router.GET("/api/login-id", isAuthenticated(), func(c *gin.Context) {common.GetLoginIdBySession(c)})

	//HTTPSサーバーを起動LSプロトコル使用※ハンドラの登録後に実行登録後に実行
	//第1引数にはポート番号 ":8080" 、第2引数にはTLS証明書のパス、第3引数には秘密鍵のパス
	// router.RunTLS(":8080", "../../certificate/localhost.crt", "../../certificate/localhost.key")

	//PORT環境変数で定義
	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
	} else {
		log.Printf("PORT is %s", port)
	}
	log.Printf("Listening on port %s.", port)
	//HTTPサーバーを起動
	// router.Run(":8080")
	// HTTPサーバーを起動し、エラーログを出力
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}

	return router
}

//ログイン中かどうかを判定するミドルウェア
//このハンドラ関数はクライアントのリクエストが処理される前に実行
func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		//セッションから取得
		id, err := redis.GetSession(c, cookieKey)
		if err != nil {
			log.Println("セッションからIDの取得に失敗しました。" , err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// セッションにログイン情報が保存されているかどうかをチェックする
		if id == nil {
			fmt.Println("セッションにユーザーIDが存在していません")
			// ログインしていない場合はログイン画面にリダイレクトする
			// c.Redirect(http.StatusFound, "/auth/login")
			c.JSON(http.StatusFound, gin.H{"message": "status 302 fail to get session id"})
			c.Abort()
		}
		fmt.Println("success get session id")
		c.Next()
	}
}
