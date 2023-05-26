package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/redis"
	"net/http"
	"os"

	"github.com/unchain1ed/server-app/model/db"
)

// func allowOrigin(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// }

func getTop(c *gin.Context, w http.ResponseWriter) {
	fmt.Println("通過N")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "home.html", gin.H{})
	// c.JSON(http.StatusOK, gin.H{})
}

func getLogin(c *gin.Context, w http.ResponseWriter) {
	fmt.Println("通過G")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func postLogin(c *gin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("通過A")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Header("Access-Control-Allow-Origin", "http://localhost:3000")

	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := db.Login(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}


	//セッションとCookieにUserIdを登録
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.NewSession(c, cookieKey, user.UserId)

	c.HTML(http.StatusOK, "mypage.html", gin.H{"user": user})
	// c.JSON(http.StatusOK, gin.H{"user": user})
}

func getSignup(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// 新規会員登録(id,password)
func postSignup(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	id := c.PostForm("user_id")
	pw := c.PostForm("password")

	user, err := db.Signup(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{"user": user})
	// c.JSON(http.StatusOK, gin.H{"user": user})
}

func getUpdate(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	// //dbパッケージからUser型のポインタを作成
	// db := &db.User{}
	// //ポインタを使ってLoggedInを呼び出し
	// user := db.LoggedIn()

	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	user := redis.GetSession(c, cookieKey)

	c.HTML(http.StatusOK, "update.html", gin.H{"user": user})
}

// 会員情報編集(id,password)
func postUpdate(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	id := c.PostForm("user_id")
	pw := c.PostForm("password")

	user, err := db.Update(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/update")
		return
	}

	//セッションとCookieにIDを登録
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.NewSession(c, cookieKey, user.UserId)

	c.Redirect(http.StatusFound, "/")
	// c.HTML(http.StatusOK, "login.html", gin.H{"user": user})
}

// マイページ画面
func getMypage(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	user := db.User{}

	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	UserId := redis.GetSession(c, cookieKey)

	if UserId != nil {
		user = db.GetOneUser(UserId.(string))
	}

	c.HTML(http.StatusOK, "mypage.html", gin.H{"user": user})

}

// ログアウト処理
func getLogout(c *gin.Context, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.DeleteSession(c, cookieKey)

	c.Redirect(http.StatusFound, "/")
}
