package login

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/unchain1ed/webapp/controller/login"
)

func TestGetLogin(t *testing.T) {
	// Ginのコンテキストを作成
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// テスト用の前提条件を設定
	// ここで必要なセットアップを行ってください。

	// GetLogin関数を呼び出す
	login.GetLogin(c)

	// レスポンスのステータスコードを検証
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスのJSONデータを検証
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 期待される結果を検証
	// responseを使って、GetLogin関数の出力を検証してください。
	// 例: assert.Equal(t, expectedValue, response["key"])
}


func TestPostLogin(t *testing.T) {
	// Ginのコンテキストを作成
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// テスト用の前提条件を設定
	// ここで必要なセットアップを行ってください。

	// GetLogin関数を呼び出す
	login.PostLogin(c)

	// レスポンスのステータスコードを検証
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスのJSONデータを検証
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 期待される結果を検証
	// responseを使って、GetLogin関数の出力を検証してください。
	// 例: assert.Equal(t, expectedValue, response["key"])
}
