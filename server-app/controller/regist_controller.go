package controller

import (
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/db"
)

func getRegist(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// 新規会員登録(id,password)
func postRegist(c *gin.Context) {
	//フォームの値を取得
	id := c.PostForm("userId")
	pw := c.PostForm("password")

	user, err := db.Signup(id, pw)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}

	log.Printf("Success user in RegisterView from DB :user %+v", user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}