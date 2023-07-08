package controller

import (
	// "time"
	"fmt"
	"net/http"
	"github.com/unchain1ed/server-app/model/redis"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	//静的ファイル
	router.LoadHTMLGlob("../../view/*.html")
	
	//クロスオリジンリソース共有を有効化
	// router.Use(cors.Default()) 

		// CORS設定
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://localhost:3000"}
		config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
		config.AllowHeaders = []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			// "Cookie",
		}
		config.AllowCredentials = true
		router.Use(cors.New(config))

	//ホーム画面
	router.GET("/", func(c *gin.Context) {getTop(c)})

	// router.POST("/blog/post", func(c *gin.Context) {postBlog(c)})

	//loginCheckGroupで/mypageと/logoutのルートパスをグループ化し、ログインチェック実施
	//ログインされていない場合はリダイレクト、ログインしている場合はそれぞれのハンドラ関数を呼び出し
	loginCheckGroup := router.Group("/", checkLogin())
	{
		//ログイン画面
		loginCheckGroup.GET("/mypage", func(c *gin.Context) {getMypage(c)})
		loginCheckGroup.GET("/logout", func(c *gin.Context) {getLogout(c)})
	}

	//ログアウトされている場合はそれぞれのハンドラ関数を呼び出し、ログインしている場合はリダイレクト
	logoutCheckGroup := router.Group("/", checkLogout())
	{
		//ログイン画面
		logoutCheckGroup.GET("/login", func(c *gin.Context) {getLogin(c)})
		logoutCheckGroup.POST("/login", func(c *gin.Context) {postLogin(c)})
		// logoutCheckGroup.POST("/login", func(c *gin.Context) {postLogin(c, c.Writer, c.Request)})
		//サインアップ画面
		logoutCheckGroup.GET("/signup", func(c *gin.Context) {getSignup(c)})
		logoutCheckGroup.POST("/signup", func(c *gin.Context) {postSignup(c)})
		//会員情報編集画面
		logoutCheckGroup.GET("/update", func(c *gin.Context) {getUpdate(c)})
		logoutCheckGroup.POST("/update", func(c *gin.Context) {postUpdate(c)})
		//ブログ記事作成画面
		logoutCheckGroup.POST("/blog/post", func(c *gin.Context) {postBlog(c)})
		//BlogOverview画面
		logoutCheckGroup.GET("/blog/overview", func(c *gin.Context) {getBlogOverview(c)})
		//BlogIDによるView画面
		logoutCheckGroup.GET("/blog/overview/post/:id", func(c *gin.Context) {getBlogViewById(c)})
		//ブログ記事編集API
		logoutCheckGroup.POST("/blog/edit", func(c *gin.Context) {editBlog(c)})
	}

	//HTTPSサーバーを起動LSプロトコル使用※ハンドラの登録後に実行登録後に実行**TODO**
	//第1引数にはポート番号 ":8080" 、第2引数にはTLS証明書のパス、第3引数には秘密鍵のパス
	// router.RunTLS(":8080", "../../certificate/server.pem", "../../certificate/server.key")

	//サーバーを起動
	router.Run(":8080")

	return router
}

//セッションにユーザーIDが存在するかチェック
//このハンドラ関数はクライアントのリクエストが処理される前に実行
//ログアウト状態であればリダイレクトし、ログイン状態であれば処理を継続
func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		//環境変数設定
		envErr := godotenv.Load("../../build/app/.env")
    	if envErr != nil {
        fmt.Println("Error loading .env file", envErr)
    	}
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		//セッションから取得
		id := redis.GetSession(c, cookieKey)

		if id == nil {
			//リダイレクト処理
			//セッションにユーザーIDが存在しない場合中断
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			} else {
			//セッションにユーザーIDが存在する場合ハンドラ関数に処理を渡します
			c.Next()
			}
	}
}

//ログイン状態であればリダイレクトし、ログアウト状態であれば処理を継続
func checkLogout() gin.HandlerFunc {
	return func(c *gin.Context) {

		//環境変数設定
		envErr := godotenv.Load("../../build/app/.env")
    	if envErr != nil {
        fmt.Println("Error loading .env file", envErr)
    	}
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		//セッションから取得
		id := redis.GetSession(c, cookieKey)

		if id != nil {
			//リダイレクト処理
			//セッションにユーザーIDが存在する場合中断
			fmt.Println("セッションにユーザーIDが存在リダイレクトTo /mypage")
			// c.Redirect(http.StatusFound, "/mypage")
			// c.Abort()
			c.Next()
		} else {
			//セッションにユーザーIDが存在しない場合ハンドラ関数に処理を渡します
			c.Next()
		}
	}	
} 

//セッションからログインIDを取得
func getLoginIdBySession(c *gin.Context) {

	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	//セッションから取得
	id := redis.GetSession(c, cookieKey)

	c.JSON(http.StatusOK, gin.H{"id": id})
}
