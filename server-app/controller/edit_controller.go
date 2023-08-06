package controller

import (
	"github.com/unchain1ed/server-app/model/entity"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
		
	// user, err := db.CheckUser(id, pw)
	// if err != nil {
	// 	c.Redirect(http.StatusMovedPermanently, "/login")
	// 	return
	// }
	fmt.Println("ID"+blogPost.ID)
	fmt.Println("loginID"+blogPost.LoginID)
	fmt.Println("title"+blogPost.Title)
	fmt.Println("content"+blogPost.Content)

	//DBにブログ記事内容を登録
	blog, err := db.Edit(blogPost.ID, blogPost.LoginID, blogPost.Title, blogPost.Content)
	if err != nil {
		log.Println("error", err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error in db.Edit": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog, "url": "/"})
}