package controller

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/unchain1ed/server-app/model/db"
)

func postEditBlog(c *gin.Context) {
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
	fmt.Println("ID"+blogPost.ID)
	fmt.Println("loginID"+blogPost.LoginID)
	fmt.Println("title"+blogPost.Title)
	fmt.Println("content"+blogPost.Content)

	//DBにブログ記事内容を登録
	blog, err := db.Edit(blogPost.ID, blogPost.LoginID, blogPost.Title, blogPost.Content)
	if err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error", err.Error());
		return
	}


	c.JSON(http.StatusOK, gin.H{"blog": blog, "url": "/blog/overview"})
}