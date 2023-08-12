package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/unchain1ed/server-app/model/entity"
	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/redis"
	"github.com/unchain1ed/server-app/model/db"
)

func getTop(c *gin.Context) {
	//セッションからloginIDを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	id, err := redis.GetSession(c, cookieKey)
	if err != nil {
		log.Printf("セッションからIDの取得に失敗しました。" , err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Get LoginId in TopView from Session :id", id); 

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// マイページ画面
func getMypage(c *gin.Context) {
	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	user, err := redis.GetSession(c, cookieKey)
	if err != nil {
		log.Printf("セッションからIDの取得に失敗しました。" , err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	//DBにブログ記事内容を登録
	blog, err := db.Create(blogPost.LoginID, blogPost.Title, blogPost.Content)
	if err != nil {
		log.Println("error", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}

// BlogOverview画面
func getBlogOverview(c *gin.Context) {
	blogs := []entity.Blog{}
	//DBからブログ情報を取得
	var err error
	blogs, err = db.GetBlogOverview()
	if err != nil {
		log.Printf("Error in GetBlogOverview of getBlogOverview ブログ概要画面でDBからブログ情報の取得に失敗しました。err: %v", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error in db.GetBlogOverview of getBlogOverview": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

// BlogView IDによる画面
func getBlogViewById(c *gin.Context) {
	id := c.Param("id")

	blog, err := db.GetBlogViewInfoById(id)
	if err != nil {
		log.Printf("Error in GetBlogViewInfoById of getBlogViewById ブログ概要画面でDBからIDの取得に失敗しました。err: %v", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error in db.GetBlogViewInfoById of getBlogViewById": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog})
}