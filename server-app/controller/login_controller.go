package controller

import (
	"errors"
	"log"
	"os"
	"net/http"

	"github.com/unchain1ed/server-app/service"
	"github.com/gin-gonic/gin"
	"github.com/unchain1ed/server-app/model/entity"
	"github.com/unchain1ed/server-app/model/redis"
	"github.com/unchain1ed/server-app/model/db"
)

func getLogin(c *gin.Context) {
	user := entity.User{}
	//セッションからuserを取得
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	UserId := redis.GetSession(c, cookieKey)

	if UserId != nil {
		user = db.GetOneUser(UserId.(string))
	}

	log.Printf("Get user in LoginView from DB :user %+v", user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func postLogin(c *gin.Context) {
	//構造体をインスタンス化
	loginUser := entity.FormUser{}
	//JSONデータのリクエストボディを構造体にバインドしてバリデーションを実行
	err := c.ShouldBindJSON(&loginUser);
	if err != nil {
		//バリデーションチェックを実行
		validationCheck := service.ValidationCheck(c, err);
		if validationCheck == false {
			err := errors.New("Error in ValidationCheck")
			log.Printf("ログイン画面リクエストJSON形式で構造体にバインドを失敗しました。loginUser.UserId: %s, err: %v", loginUser.UserId, err.Error());
			c.JSON(http.StatusBadRequest, gin.H{"error in c.ShouldBindJSON of postLogin": err.Error()})
			return
		}
	}

	//DB上の会員情報と照合チェック
	user, err := db.CheckUser(loginUser.UserId, loginUser.Password)
	if err != nil {
		err := errors.New("Error in CheckUser")
		log.Printf("ログイン画面DB上の会員情報と照合に失敗しました。loginUser.UserId: %s, err: %v", loginUser.UserId, err.Error());
		c.JSON(http.StatusBadRequest, gin.H{"error in db.CheckUser of postLogin": err.Error()})
		return
	}

	//セッションとCookieにUserIdを登録
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	redis.NewSession(c, cookieKey, user.UserId)

	c.JSON(http.StatusOK, gin.H{"user": user})
}