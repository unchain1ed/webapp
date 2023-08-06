package controller

import (
	"errors"
	"github.com/unchain1ed/server-app/service"
	"github.com/unchain1ed/server-app/model/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/db"
)

// 新規会員登録(id,password)
func postRegist(c *gin.Context) {
	//構造体をインスタンス化
	registUser := entity.FormUser{}
	//JSONデータのリクエストボディを構造体にバインドしてバリデーションを実行
	err := c.ShouldBindJSON(&registUser);
	if err != nil {
		//バリデーションチェックを実行
		validationCheck := service.ValidationCheck(c, err);
		if validationCheck == false {
			err := errors.New("Error in ValidationCheck")
			log.Printf("会員登録画面リクエストJSON形式で構造体にバインドを失敗しました。registUser.UserId: %s, err: %v", registUser.UserId, err.Error());
			c.JSON(http.StatusBadRequest, gin.H{"error in c.ShouldBindJSON of postRegist": err.Error()})
			return
		}
	}

	user, err := db.Signup(registUser.UserId, registUser.Password)
	if err != nil {
		err := errors.New("Error in db.Signup")
			log.Printf("DBに会員情報の登録に失敗しました。registUser.UserId: %s, err: %v", registUser.UserId, err.Error());
			c.JSON(http.StatusBadRequest, gin.H{"error in db.Signup of postRegist": err.Error()})
		return
	}

	log.Printf("Success user in RegisterView from DB :user %+v", user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}