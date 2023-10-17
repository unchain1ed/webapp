package common

import (
	"os"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/unchain1ed/webapp/controller/common"
	"net/http"
	"net/http/httptest"
	"github.com/unchain1ed/webapp/model/redis"
	"github.com/gin-gonic/gin"
	"testing"
)


// func TestGetLoginIdBySession(t *testing.T) {
// 	// Arrange ---
// 	response := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(response)
// 	c.Request, _ = http.NewRequest(
// 		http.MethodGet,
// 		"/api/login-id",
// 		nil,
// 	)
 
// 	redis.NewSession(c, "loginUserIdKey","root")

// 	// Act ---
// 	common.GetLoginIdBySession(c)
 
// 	// Assert ---
// 	// var product products.Product
// 	// err := json.Unmarshal(response.Body.Bytes(), "root")
// 	assert.EqualValues(t, http.StatusOK, response.Code)
// 	// assert.Nil(t, err)
// 	// fmt.Println(product)
// 	// assert.EqualValues(t, uint64(123), product.ID)
// }

func TestGetLoginIdBySession(t *testing.T) {
    // Arrange ---
    response := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(response)
    c.Request, _ = http.NewRequest(
        http.MethodGet,
        "/api/login-id",
        nil,
    )

    // セッション情報を設定
    os.Setenv("LOGIN_USER_ID_KEY", "loginUserIdKey")  // テスト用の環境変数を設定
    redis.NewSession(c, "loginUserIdKey", "root")

    // Act ---
    common.GetLoginIdBySession(c)

    // Assert ---
    assert.Equal(t, http.StatusOK, response.Code) // ステータスコードが期待通りか確認

    // レスポンスボディをJSONとしてパース
    var responseJSON map[string]interface{}
    err := json.Unmarshal(response.Body.Bytes(), &responseJSON)
    if err != nil {
        t.Fatalf("Failed to parse JSON response: %v", err)
    }

    // レスポンスボディの"id"フィールドを確認
    expectedID := "root"  // 期待されるIDを設定
    if id, found := responseJSON["id"]; found {
        assert.Equal(t, expectedID, id)
    } else {
        t.Fatalf("Response does not contain 'id' field")
    }
}
