package controller

import (
	"github.com/unchain1ed/server-app/model/redis"
	"errors"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/unchain1ed/server-app/model/db"
	"github.com/go-playground/validator/v10"
)

type User struct {
	NowId     string `json:"nowId" binding:"required,min=2,max=10"`
	ChangeId     string `json:"changeId" binding:"required,min=2,max=10"`
	// Password string `json:"password" binding:"required,min=4"`
}

// type BlogPost struct {
// 	ID string `json:"id"`
// 	LoginID string `json:"loginID"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// 会員情報編集(id)
func postSettingId(c *gin.Context) {
	var user User

	// JSONデータをUser構造体にバインドしてバリデーションを実行
	if err := c.ShouldBindJSON(&user); err != nil {
		// バリデーションエラーが発生した場合はエラーレスポンスを返す
		var verr validator.ValidationErrors
		if ok := errors.As(err, &verr); ok {
			var errors []string
			for _, e := range verr {
				errors = append(errors, fmt.Sprintf("%s validation failed on the %s field", e.Tag(), e.Field()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("リクエストID"+user.NowId+user.ChangeId)

	//DBにIDを変更
	blog, err := db.UpdateId(user.ChangeId, user.NowId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("error", err.Error());
		return
	}

	//redisでセッション破棄、新IDでセッション作成
	err = redis.UpdateSession(c, user.ChangeId, user.NowId)
	if err != nil {
		log.Println("セッション破棄と変更後IDでセッション作成に失敗しました。",err)
	}

	c.JSON(http.StatusOK, gin.H{"blog": blog, "url": "/"})
}

// // 会員情報編集(password)
// func postUpdateid(c *gin.Context) {
// 	id := c.PostForm("userId")
// 	pw := c.PostForm("password")

// 	user, err := db.Update(id, pw)
// 	if err != nil {
// 		c.Redirect(http.StatusMovedPermanently, "/update")
// 		return
// 	}

// 	//セッションとCookieにIDを登録
// 	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
// 	redis.NewSession(c, cookieKey, user.UserId)

// 	c.Redirect(http.StatusFound, "/")
// }