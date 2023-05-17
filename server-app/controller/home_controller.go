package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/unchain1ed/server-app/model/db"
)


// func CORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// クロスオリジン用にセット
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
// 		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		// w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// w.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,UPDATE,OPTIONS")
//         // w.Header().Set("Content-Type", "application/json")
// 	})
// }

// func init(c *gin.Context) {
// 	c.http.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// }

func getTop(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func getLogin(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func postLogin(c *gin.Context,w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := db.Login(id, pw)
	if err != nil {
		c.Redirect(301, "/login")
		return
	}
	c.HTML(http.StatusOK, "blog.html", gin.H{"user": user})

	//ログイン成功後のリダイレクト処理
	// http.Redirect(w, r, "/login", http.StatusFound)

	
}

func getSignup(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

//新規会員登録(id,password)
func postSignup(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	id := c.PostForm("user_id")
	pw := c.PostForm("password")

	user, err := db.Signup(id, pw)
	if err != nil {
		c.Redirect(301, "/signup")
		return
	}
	c.HTML(http.StatusOK, "signup.html", gin.H{"user": user})
}

func getUpdate(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	//dbパッケージからUser型のポインタを作成
	db := &db.User{}
	//ポインタを使ってLoggedInを呼び出し
	user := db.LoggedIn()

	c.HTML(http.StatusOK, "update.html", gin.H{"user": user})
}

//会員情報編集(id,password)
func postUpdate(c *gin.Context,w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	id := c.PostForm("user_id")
	pw := c.PostForm("password")

	user, err := db.Update(id, pw)
	if err != nil {
		c.Redirect(301, "/update")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"user": user})
}
