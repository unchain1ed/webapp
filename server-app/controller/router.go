package controller

import (
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
	// クロスオリジンリソース共有を有効化
	router.Use(cors.Default()) 

	//静的ファイル
	router.LoadHTMLGlob("../../view/*.html")

	//ホーム画面
	router.GET("/", func(c *gin.Context) {getTop(c, c.Writer)})
	


	// router.GET("/login", func(c *gin.Context) {getLogin(c, c.Writer)})

	

	//loginCheckGroupで/mypageと/logoutのルートパスをグループ化し、ログインチェック実施
	//ログインされていない場合はリダイレクト、ログインしている場合はそれぞれのハンドラ関数を呼び出し
	loginCheckGroup := router.Group("/", checkLogin())
	{

		//ログイン画面
		loginCheckGroup.GET("/mypage", func(c *gin.Context) {getMypage(c, c.Writer)})
		loginCheckGroup.GET("/logout", func(c *gin.Context) {getLogout(c, c.Writer)})
	}

	//ログアウトされている場合はそれぞれのハンドラ関数を呼び出し、ログインしている場合はリダイレクト
	logoutCheckGroup := router.Group("/", checkLogout())
	{
		//ログイン画面
		logoutCheckGroup.GET("/login", func(c *gin.Context) {getLogin(c, c.Writer)})
		logoutCheckGroup.POST("/login", func(c *gin.Context) {postLogin(c, c.Writer, c.Request)})
		//サインアップ画面
		logoutCheckGroup.GET("/signup", func(c *gin.Context) {getSignup(c, c.Writer)})
		logoutCheckGroup.POST("/signup", func(c *gin.Context) {postSignup(c, c.Writer)})
		//会員情報編集画面
		logoutCheckGroup.GET("/update", func(c *gin.Context) {getUpdate(c, c.Writer)})
		logoutCheckGroup.POST("/update", func(c *gin.Context) {postUpdate(c, c.Writer)})
	}

	//HTTPSサーバーを起動LSプロトコル使用※ハンドラの登録後に実行
	//第1引数にはポート番号 ":8080" 、第2引数にはTLS証明書のパス、第3引数には秘密鍵のパス
	router.RunTLS(":8080", "../../certificate/server.pem", "../../certificate/server.key")

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
			c.Redirect(http.StatusFound, "/mypage")
			c.Abort()
		} else {
			//セッションにユーザーIDが存在しない場合ハンドラ関数に処理を渡します
			c.Next()
		}
	}	
} 

