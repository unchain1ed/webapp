package controller

import (
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/unchain1ed/server-app/model/redis"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func init() {
	//環境変数設定
	//パッケージがインポートされる際に一度だけ実行される初期化処理
	envErr := godotenv.Load("../../build/app/.env")
	if envErr != nil {
		fmt.Println("Error loading .env file", envErr)
	}
}	

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

	//ホーム画面
	router.GET("/", isAuthenticated(), func(c *gin.Context) {getTop(c)})

	//***ログイン画面***
	//ログイン画面
	router.GET("/login", func(c *gin.Context) {getLogin(c)})
	router.POST("/login", func(c *gin.Context) {postLogin(c)})

	//***ブログ概要画面***
	//ブログ記事作成画面
	router.POST("/blog/post", isAuthenticated(), func(c *gin.Context) {postBlog(c)})
	//BlogOverview画面
	router.GET("/blog/overview", isAuthenticated(),func(c *gin.Context) {getBlogOverview(c)})
	//BlogIDによるView画面
	router.GET("/blog/overview/post/:id", isAuthenticated(), func(c *gin.Context) {getBlogViewById(c)})
	//ブログ記事編集API
	router.POST("/blog/edit", isAuthenticated(), func(c *gin.Context) {postEditBlog(c)})
	//ブログ記事消去API
	router.GET("/blog/delete/:id", isAuthenticated(), func(c *gin.Context) {getDeleteBlog(c)})

	//***会員情報編集画面***
	//ID編集API
	router.POST("/update/id", isAuthenticated(), func(c *gin.Context) {postSettingId(c)})

	//***ログアウト画面***
	//ログアウト実行API
	router.POST("/logout", isAuthenticated(), func(c *gin.Context) {decideLogout(c)})

	//***会員情報登録画面***
	//登録画面遷移
	router.POST("/regist", isAuthenticated(), func(c *gin.Context) {postRegist(c)})

	//***共通API***
	//セッションからログインIDを取得するAPI
	router.GET("/api/login-id", isAuthenticated(), func(c *gin.Context) {getLoginIdBySession(c)})

	//HTTPSサーバーを起動LSプロトコル使用※ハンドラの登録後に実行登録後に実行
	//第1引数にはポート番号 ":8080" 、第2引数にはTLS証明書のパス、第3引数には秘密鍵のパス
	// router.RunTLS(":8080", "../../certificate/localhost.crt", "../../certificate/localhost.key")

	//HTTPサーバーを起動
	router.Run(":8080")

	return router
}

//ログイン中かどうかを判定するミドルウェア
//このハンドラ関数はクライアントのリクエストが処理される前に実行
func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("通過isAuthenticated")
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		//セッションから取得
		id, err := redis.GetSession(c, cookieKey)
		if err != nil {
			log.Printf("セッションからIDの取得に失敗しました。" , err.Error())
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
