package redis

import (
	"os"
	"net/http"
	"log"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var conn *redis.Client

func init() {
	var dbHost string
	// Dockerコンテナ内での接続先を指定
	if os.Getenv("DOCKER_ENV") == "true" {
		dbHost = "redis:6379"
	} else {
		// ローカル環境での接続先を指定
		dbHost = "localhost:6379"
	}
	//Redisデータベース接続のためRedisクライアント作成
	conn = redis.NewClient(&redis.Options{
		Addr:     dbHost,
		Password: "",
		DB:       0,
	})
}

func NewSession(c *gin.Context, cookieKey, redisValue string) {
	slice := make([]byte, 64)
	//ランダムなバイト列を生成
	if _, err := io.ReadFull(rand.Reader, slice); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}

	//バイト配列を base64 エンコードして文字列に変換
	newRedisKey := base64.URLEncoding.EncodeToString(slice)

	//Redisにセッションを登録
	if err := conn.Set(c, newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}
	fmt.Println("HTTPレスポンスヘッダcookieKey"+cookieKey)
	fmt.Println("HTTPレスポンスヘッダnewRedisKey"+newRedisKey)
	//HTTPレスポンスヘッダーにCookieを設定
	cookie := &http.Cookie{
		Name:     cookieKey,
		Value:    newRedisKey,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   0,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	// クッキーをレスポンスに設定
	http.SetCookie(c.Writer, cookie)
}

func GetSession(c *gin.Context, cookieKey string) interface{} {
	fmt.Println("通過cookieKey"+cookieKey)

	//クライアントのリクエストに含まれるセッションのクッキー値を取得
	redisKey, err := c.Cookie(cookieKey)
	if err != nil {
		fmt.Println("セッションのクッキーが見つかりませんでした。,err："+err.Error())
	}

	fmt.Println("通過redisKey"+redisKey)

	//取得したセッションのクッキー値を使用して、Redisから対応するセッションデータを取得
	redisValue, err := conn.Get(c, redisKey).Result()
	fmt.Println("通過redisValue"+redisValue)
	log.Println("redisKey,redisValue,err："+redisKey,redisValue,err)
	
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません。")
		return nil
	case err != nil:
		fmt.Println("Session取得時にエラー発生：" + err.Error())
		return nil
	}
	return redisValue
}

func DeleteSession(c *gin.Context, cookieKey string) {
	redisId, err := c.Cookie(cookieKey)
	if err != nil {
		fmt.Println("セッションのクッキーが見つかりませんでした。")
		return
	}
	//Redisからセッションを削除
	cmd := conn.Del(c, redisId)
	if cmd.Val() == 0 {
		fmt.Println("Redisからセッションを削除できませんでした。")
	} else {
		fmt.Println("Redisからセッションを削除しました。：", cmd.String())
	}
	
	//クライアントのブラウザに保存されているセッションのクッキーを削除
	cookie := &http.Cookie{
		Name:     cookieKey,
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:    -1,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	// クッキーをレスポンスに設定
	http.SetCookie(c.Writer, cookie)
}