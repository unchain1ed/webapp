package controller

import (
	"github.com/joho/godotenv"
	"log"
	"gorm.io/gorm"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/redis"
	"net/http"
	"os"

	"github.com/unchain1ed/server-app/model/db"
)

type BlogPost struct {
	LoginID string `json:"loginID"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Blog struct {
	gorm.Model //共通カラム
	LoginID string
	Title string
	Content string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func getTop(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "text/plain")
	//環境変数設定
	envErr := godotenv.Load("../../build/app/.env")
	if envErr != nil {
		fmt.Println("Error loading .env file", envErr)
	}
	//セッションからloginIDを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	loginID := redis.GetSession(c, cookieKey)
fmt.Println(loginID)
	c.JSON(http.StatusOK, gin.H{"loginID": loginID})
}

func getLogin(c *gin.Context) {
	user := db.User{}
	// //セッションからuserを取得
	// cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	// user := redis.GetSession(c, cookieKey)

	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	UserId := redis.GetSession(c, cookieKey)

	if UserId != nil {
		user = db.GetOneUser(UserId.(string))
	}
	// fmt.Println("UserId"+UserId.(string))
	// fmt.Println("user.UserId"+user.UserId)
	// c.HTML(http.StatusOK, "login.html", gin.H{})
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func postLogin(c *gin.Context) {
	//フォームの値を取得
	id := c.PostForm("userId")
	pw := c.PostForm("password")
		
	user, err := db.Login(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	fmt.Println("user.UserId"+user.UserId)
	fmt.Println("id password"+id +pw)

	//セッションとCookieにUserIdを登録
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.NewSession(c, cookieKey, user.UserId)

	// c.HTML(http.StatusOK, "mypage.html", gin.H{"user": user})
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// 新規会員登録(id,password)
func postSignup(c *gin.Context) {
	id := c.PostForm("userId")
	pw := c.PostForm("password")

	user, err := db.Signup(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{"user": user})
	// c.JSON(http.StatusOK, gin.H{"user": user})
}

func getUpdate(c *gin.Context) {
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
func postUpdate(c *gin.Context) {
	id := c.PostForm("userId")
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
func getMypage(c *gin.Context) {
	// user := db.User{}

	// //セッションからuserを取得
	// cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	// UserId := redis.GetSession(c, cookieKey)

	// if UserId != nil {
	// 	user = db.GetOneUser(UserId.(string))
	// }
	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	user := redis.GetSession(c, cookieKey)

	// c.HTML(http.StatusOK, "login.html", gin.H{})
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// ログアウト処理
func getLogout(c *gin.Context) {
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.DeleteSession(c, cookieKey)

	c.Redirect(http.StatusFound, "/login")
}

func postBlog(c *gin.Context) {
	// c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Content-Type", "text/plain")

	// JSON形式のリクエストボディを構造体にバインドする
	var blogPost BlogPost
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		
	// user, err := db.Login(id, pw)
	// if err != nil {
	// 	c.Redirect(http.StatusMovedPermanently, "/login")
	// 	return
	// }
	fmt.Println("loginID"+blogPost.LoginID)
	fmt.Println("title"+blogPost.Title)
	fmt.Println("content"+blogPost.Content)

	//DBにブログ記事内容を登録
	blog, err := db.Create(blogPost.LoginID, blogPost.Title, blogPost.Content)
	if err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error", err.Error());
		return
	}


	c.JSON(http.StatusOK, gin.H{"blog": blog, "url": "/blog/overview"})
}

// BlogOverview画面
func getBlogOverview(c *gin.Context) {
	// blog := db.Blog{}
	var blogs []db.Blog

	// //セッションからuserを取得
	// cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	// UserId := redis.GetSession(c, cookieKey)

	blogs = db.GetBlogOverview()

	//セッションからuserを取得


	// c.HTML(http.StatusOK, "login.html", gin.H{})
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

// BlogView IDによる画面
func getBlogViewById(c *gin.Context) {
	
	//  var blog = Blog{}
	//  blog := Blog{}

	//  blog := &Blog{}
	// var blogs []db.Blog

	id := c.Param("id")
	fmt.Println("Param"+id)
	// //セッションからuserを取得
	// cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	// UserId := redis.GetSession(c, cookieKey)

	blog ,err := db.GetBlogViewInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//セッションからuserを取得


	c.JSON(http.StatusOK, gin.H{"blog": blog})
}