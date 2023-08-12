package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/db"
	"github.com/unchain1ed/server-app/service"
	"github.com/unchain1ed/server-app/model/entity"
	"github.com/unchain1ed/server-app/model/redis"
)

// 会員情報編集(id)
func postSettingId(c *gin.Context) {
	user := entity.UserChange{}

	//リクエストをGo構造体にバインド
	err := c.ShouldBindJSON(&user);
	//JSONデータをUser構造体にバインドしてバリデーションを実行
	if err != nil {
		//バリデーションチェックを実行
		err := service.ValidationCheck(c, err);
		if err != nil {
		
			log.Printf("セッションidが一致しませんでした。user.ChangeId: %s, user.NowId: %s, err: %v", user.ChangeId, user.NowId, err.Error());
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

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
// err = redis.NewSession(c, cookieKey, user.UserId)
// if err != nil {
// 	log.Printf("Error in NewSession ログイン画面DB上のセッションにIDの登録に失敗しました。loginUser.UserId: %s, err: %v", loginUser.UserId, err.Error());
// 	c.JSON(http.StatusBadRequest, gin.H{"error in redis.NewSession of postLogin": err.Error()})
// 	return
// }

// 	c.Redirect(http.StatusFound, "/")
// }