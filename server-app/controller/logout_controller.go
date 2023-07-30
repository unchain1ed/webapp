package controller

import (
	"log"
	"github.com/unchain1ed/server-app/model/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ログアウト処理
func decideLogout(c *gin.Context) {

	id := c.PostForm("userId")

	err := redis.DeleteSession(c, id)
	if err != nil {
		log.Println("セッションを消去できませんでした。,err："+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error in decideLogout": err.Error()})
		return
	}
	log.Println("Success Logout :id", id); 

	c.JSON(http.StatusOK, gin.H{"Success Logout": "auth/login"})
}