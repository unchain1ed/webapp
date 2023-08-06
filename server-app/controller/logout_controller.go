package controller

import (
	"github.com/unchain1ed/server-app/model/db"
	"errors"
	"github.com/unchain1ed/server-app/service"
	"github.com/unchain1ed/server-app/model/entity"
	"log"
	"github.com/unchain1ed/server-app/model/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ログアウト処理
func decideLogout(c *gin.Context) {
	//構造体をインスタンス化
	logoutUser := entity.FormUser{}
	//JSONデータのリクエストボディを構造体にバインドしてバリデーションを実行
	err := c.ShouldBindJSON(&logoutUser);
	if err != nil {
		//バリデーションチェックを実行
		validationCheck := service.ValidationCheck(c, err);
		if validationCheck == false {
			err := errors.New("Error in ValidationCheck")
			log.Printf("ログアウト画面リクエストJSON形式で構造体にバインドを失敗しました。logoutUser.UserId: %s, err: %v", logoutUser.UserId, err.Error());
			c.JSON(http.StatusBadRequest, gin.H{"error in c.ShouldBindJSON of decideLogout": err.Error()})
			return
		}
	}

	//DB上の会員情報と照合チェック
	user, err := db.CheckUser(logoutUser.UserId, logoutUser.Password)
	if err != nil {
		log.Printf("ログアウト画面DB上の会員情報と照合に失敗しました。logoutUser.UserId: %s", logoutUser.UserId);
		c.JSON(http.StatusBadRequest, gin.H{"error in db.CheckUser of decideLogout": err.Error()})
		return
	}

	//Redisよりログイン情報セッションを消去
	err = redis.DeleteSession(c, user.UserId)
	if err != nil {
		log.Println("セッションを消去できませんでした。,err："+err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error in decideLogout": err.Error()})
		return
	}
	log.Println("ログアウトに成功しました。Success Logout :id", logoutUser.UserId); 

	c.JSON(http.StatusOK, gin.H{"Success Logout": "auth/login"})
}