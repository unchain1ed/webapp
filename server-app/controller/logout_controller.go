package controller

import (
	"log"
	"github.com/unchain1ed/server-app/model/redis"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ログアウト処理
func decideLogout(c *gin.Context) {

	id := c.PostForm("userId")

	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.DeleteSession(c, cookieKey, id)

	log.Println("Success Logout :id", id); 

	c.JSON(http.StatusOK, gin.H{"Success Logout": "auth/login"})
}