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

func init() {
    // このパッケージがインポートされる際に一度だけ実行される初期化処理
	//環境変数設定
	envErr := godotenv.Load("../../build/app/.env")
	if envErr != nil {
		fmt.Println("Error loading .env file", envErr)
	}
}	

func GetRouter() *gin.Engine {
	router := gin.Default()

	//静的ファイル
	router.LoadHTMLGlob("../../view/*.html")
	
	//クロスオリジンリソース共有を有効化
	// router.Use(cors.Default()) 

		// CORS設定
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
		router.Use(cors.New(config))

	//ホーム画面
	router.GET("/", isAuthenticated(), func(c *gin.Context) {getTop(c)})

	//loginCheckGroupでルートパスをグループ化し、ログインチェック実施
	//ログインされていない場合はリダイレクト、ログインしている場合はそれぞれのハンドラ関数を呼び出し
	//【ログイン中】
	// loginCheckGroup := router.Group("/", checkLogin())
	// {
	// 	// //ログイン画面
	// 	// loginCheckGroup.GET("/mypage", func(c *gin.Context) {getMypage(c)})
	// 	// loginCheckGroup.GET("/logout", func(c *gin.Context) {getLogout(c)})
	// 	// //ブログ記事作成画面
	// 	// loginCheckGroup.POST("/blog/post", func(c *gin.Context) {postBlog(c)})
	// 	// //BlogOverview画面
	// 	// loginCheckGroup.GET("/blog/overview", func(c *gin.Context) {getBlogOverview(c)})
	// 	// //BlogIDによるView画面
	// 	// loginCheckGroup.GET("/blog/overview/post/:id", func(c *gin.Context) {getBlogViewById(c)})
	// 	// //ブログ記事編集API
	// 	// loginCheckGroup.POST("/blog/edit", func(c *gin.Context) {postEditBlog(c)})
	// 	// //ブログ記事消去API
	// 	// loginCheckGroup.GET("/blog/delete/:id", func(c *gin.Context) {getDeleteBlog(c)})
	// 	// //会員情報編集画面
	// 	// loginCheckGroup.GET("/update", func(c *gin.Context) {getUpdate(c)})
	// 	// loginCheckGroup.POST("/update", func(c *gin.Context) {postUpdate(c)})
	// }

	//***ログイン画面***
	//ログイン画面
	router.GET("/login", func(c *gin.Context) {getLogin(c)})
	router.POST("/login", func(c *gin.Context) {postLogin(c)})
	// //ログイン画面
	// router.GET("/mypage", func(c *gin.Context) {getMypage(c)})

	//***ブログ画面***
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

	//会員情報編集画面
	router.GET("/update", isAuthenticated(), func(c *gin.Context) {getUpdate(c)})
	router.POST("/update", isAuthenticated(), func(c *gin.Context) {postUpdate(c)})

	//***ログアウト画面***
	//ログアウト実行API
	router.POST("/logout", isAuthenticated(), func(c *gin.Context) {decideLogout(c)})

	//***会員情報登録画面***
	//登録画面遷移
	router.GET("/regist", isAuthenticated(), func(c *gin.Context) {getRegist(c)})
	router.POST("/regist", isAuthenticated(), func(c *gin.Context) {postRegist(c)})

	//***共通API***
	//セッションからログインIDを取得するAPI
	router.GET("/api/login-id", isAuthenticated(), func(c *gin.Context) {getLoginIdBySession(c)})





	// //ログアウトされている場合はそれぞれのハンドラ関数を呼び出し、ログインしている場合はリダイレクト
	// //【ログアウト中】
	// logoutCheckGroup := router.Group("/", checkLogout())
	// {
	// 	//ログイン画面
	// 	logoutCheckGroup.GET("/mypage", func(c *gin.Context) {getMypage(c)})
	// 	//ブログ記事作成画面
	// 	logoutCheckGroup.POST("/blog/post", func(c *gin.Context) {postBlog(c)})
	// 	//BlogOverview画面
	// 	logoutCheckGroup.GET("/blog/overview", func(c *gin.Context) {getBlogOverview(c)})
	// 	//BlogIDによるView画面
	// 	logoutCheckGroup.GET("/blog/overview/post/:id", func(c *gin.Context) {getBlogViewById(c)})
	// 	//ブログ記事編集API
	// 	logoutCheckGroup.POST("/blog/edit", func(c *gin.Context) {postEditBlog(c)})
	// 	//ブログ記事消去API
	// 	logoutCheckGroup.GET("/blog/delete/:id", func(c *gin.Context) {getDeleteBlog(c)})
	// 	//会員情報編集画面
	// 	logoutCheckGroup.GET("/update", func(c *gin.Context) {getUpdate(c)})
	// 	logoutCheckGroup.POST("/update", func(c *gin.Context) {postUpdate(c)})
	// 	//ログアウト実行API
	// 	logoutCheckGroup.POST("/logout", func(c *gin.Context) {decideLogout(c)})
	// 	//ログイン中のIDをセッションから取得
	// 	logoutCheckGroup.GET("/api/login-id", func(c *gin.Context) {getLoginIdBySession(c)})

	// 	// //ログイン画面
	// 	// logoutCheckGroup.GET("/login", func(c *gin.Context) {getLogin(c)})
	// 	// logoutCheckGroup.POST("/login", func(c *gin.Context) {postLogin(c)})
	// 	//サインアップ画面
	// 	logoutCheckGroup.GET("/signup", func(c *gin.Context) {getSignup(c)})
	// 	logoutCheckGroup.POST("/signup", func(c *gin.Context) {postSignup(c)})
	// }

	//ーーNG
	// loginCheckGroup := router.Group("/", checkLogin())
	// {
	// 	//ログイン画面
	// 	loginCheckGroup.GET("/mypage", func(c *gin.Context) {getMypage(c)})
	// 	loginCheckGroup.GET("/logout", func(c *gin.Context) {getLogout(c)})
	// 	//ブログ記事作成画面
	// 	loginCheckGroup.POST("/blog/post", func(c *gin.Context) {postBlog(c)})
	// 	//BlogOverview画面
	// 	loginCheckGroup.GET("/blog/overview", func(c *gin.Context) {getBlogOverview(c)})
	// 	//BlogIDによるView画面
	// 	loginCheckGroup.GET("/blog/overview/post/:id", func(c *gin.Context) {getBlogViewById(c)})
	// 	//ブログ記事編集API
	// 	loginCheckGroup.POST("/blog/edit", func(c *gin.Context) {postEditBlog(c)})
	// 	//ブログ記事消去API
	// 	loginCheckGroup.GET("/blog/delete/:id", func(c *gin.Context) {getDeleteBlog(c)})
	// 	//会員情報編集画面
	// 	loginCheckGroup.GET("/update", func(c *gin.Context) {getUpdate(c)})
	// 	loginCheckGroup.POST("/update", func(c *gin.Context) {postUpdate(c)})

	// 	//ログイン画面
	// 	loginCheckGroup.GET("/login", func(c *gin.Context) {getLogin(c)})
	// 	loginCheckGroup.POST("/login", func(c *gin.Context) {postLogin(c)})
	// 	//サインアップ画面
	// 	loginCheckGroup.GET("/signup", func(c *gin.Context) {getSignup(c)})
	// 	loginCheckGroup.POST("/signup", func(c *gin.Context) {postSignup(c)})
	// }


	//HTTPSサーバーを起動LSプロトコル使用※ハンドラの登録後に実行登録後に実行**TODO**
	//第1引数にはポート番号 ":8080" 、第2引数にはTLS証明書のパス、第3引数には秘密鍵のパス
	// router.RunTLS(":8080", "../../certificate/localhost.crt", "../../certificate/localhost.key")

	//HTTPサーバーを起動
	router.Run(":8080")

	return router
}

//セッションにユーザーIDが存在するかチェック
//このハンドラ関数はクライアントのリクエストが処理される前に実行
//ログアウト状態であればリダイレクトし、ログイン状態であれば処理を継続
func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// ログイン中かどうかを判定するミドルウェア
func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("通過isAuthenticated")
		cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
		//セッションから取得
		id := redis.GetSession(c, cookieKey)

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
