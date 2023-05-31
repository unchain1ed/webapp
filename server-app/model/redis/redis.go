package redis

import (
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
	//Redisデータベース接続のためRedisクライアント作成
	conn = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		// Addr:     "redis:6379", //docker起動の場合
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

	fmt.Println("通過redisValue"+redisValue)

	//Redisにセッションを登録
	if err := conn.Set(c, newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生：" + err.Error())
	}
	fmt.Println("HTTPレスポンスヘッダcookieKey"+cookieKey)
	fmt.Println("HTTPレスポンスヘッダnewRedisKey"+newRedisKey)
	//HTTPレスポンスヘッダーにCookieを設定
	// c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
	// クッキーの設定
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
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	//クライアントのリクエストに含まれるセッションのクッキー値を取得
	redisKey, _ := c.Cookie(cookieKey)

	// クライアントのリクエストに含まれるセッションのクッキー値を取得
	// 以下が今NG　redisKeyが空
    // cookie, err := c.Request.Cookie(cookieKey)


    // if err != nil {
    //     fmt.Println("Cookieの取得に失敗しました:", err)
    //     return nil
    // }
    // redisKey := cookie.Value

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
	redisId, _ := c.Cookie(cookieKey)
	fmt.Println("通過redisId"+redisId)
	//Redisからセッションを削除
	conn.Del(c, redisId)
	//クライアントのブラウザに保存されているセッションのクッキーを削除
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, true)
}