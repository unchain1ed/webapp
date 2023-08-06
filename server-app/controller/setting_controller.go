package controller

import (
	"errors"
	"github.com/unchain1ed/server-app/service"
	"github.com/unchain1ed/server-app/model/entity"
	"github.com/unchain1ed/server-app/model/redis"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/unchain1ed/server-app/model/db"
)

// 会員情報編集(id)
func postSettingId(c *gin.Context) {
	user := entity.UserChange{}

	//JSONデータをUser構造体にバインドしてバリデーションを実行
	err := c.ShouldBindJSON(&user);
	if err != nil {
		//バリデーションチェックを実行
		validationCheck := service.ValidationCheck(c, err);
		if validationCheck == false {
			err := errors.New("Error in ValidationCheck")
			log.Printf("セッションidが一致しませんでした。user.ChangeId: %s, user.NowId: %s, err: %v", user.ChangeId, user.NowId, err);
			return
		}
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