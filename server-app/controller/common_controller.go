package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/redis"
)

//セッションからログインIDを取得するAPI
func getLoginIdBySession(c *gin.Context) {
	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	//セッションから取得
	id := redis.GetSession(c, cookieKey)

	log.Println("Get LoginId bySession :id", id); 

	c.JSON(http.StatusOK, gin.H{"id": id})
}