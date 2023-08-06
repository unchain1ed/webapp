package controller

import (
	"github.com/unchain1ed/server-app/model/entity"
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/redis"
	"github.com/unchain1ed/server-app/model/db"
)

// type BlogPost struct {
// 	ID string `json:"id"`
// 	LoginID string `json:"loginID" binding:"required,min=2,max=10"`
// 	Title   string `json:"title" binding:"required,min=1,max=50"`
// 	Content string `json:"content" binding:"required,min=1,max=8000"`
// }

// type Blog struct {
// 	gorm.Model //共通カラム
// 	LoginID string
// 	Title string
// 	Content string
// 	CreatedAt string
// 	UpdatedAt string
// 	DeletedAt string
// }

func getTop(c *gin.Context) {
	//セッションからloginIDを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	id := redis.GetSession(c, cookieKey)

	log.Println("Get LoginId in TopView from Session :id", id); 

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// マイページ画面
func getMypage(c *gin.Context) {
	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	user := redis.GetSession(c, cookieKey)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func postBlog(c *gin.Context) {	
	// JSON形式のリクエストボディを構造体にバインドする
	blogPost := entity.BlogPost{}
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		log.Printf("ブログ記事作成画面でJSON形式構造体にバインドを失敗しました。" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		
	// user, err := db.CheckUser(id, pw)
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
		log.Println("error", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog, "url": "/"})
}

// BlogOverview画面
func getBlogOverview(c *gin.Context) {
	// var blogs []entity.Blog
	blogs := []entity.Blog{}

	// //セッションからuserを取得
	// cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	// UserId := redis.GetSession(c, cookieKey)

	blogs = db.GetBlogOverview()

	//セッションからuserを取得

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

// BlogView IDによる画面
func getBlogViewById(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("Param"+id)

	blog ,err := db.GetBlogViewInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}