package controller

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/entity"
	"github.com/unchain1ed/server-app/model/db"
)

func postEditBlog(c *gin.Context) {
	// JSON形式のリクエストボディを構造体にバインドする
	blogPost := entity.BlogPost{}
	if err := c.ShouldBindJSON(&blogPost); err != nil {
		log.Printf("ブログ編集画面リクエストJSON形式で構造体にバインドを失敗しました。" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error in c.ShouldBindJSON": err.Error()})
		return
	}

	//DBにブログ記事内容を登録
	blog, err := db.Edit(blogPost.ID, blogPost.LoginID, blogPost.Title, blogPost.Content)
	if err != nil {
		log.Println("error", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error in db.Edit": err.Error()})
		return
	}

	log.Printf("Success Edit Blog :blog %+v", blog)
	c.JSON(http.StatusOK, gin.H{"blog": blog})
}