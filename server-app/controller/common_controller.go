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
	id, err := redis.GetSession(c, cookieKey)
	if err != nil {
		log.Printf("セッションからIDの取得に失敗しました。" , err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Get LoginId bySession :id", id); 

	c.JSON(http.StatusOK, gin.H{"id": id})
}