package common

import (
	// "encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/unchain1ed/webapp/controller/common"
	"net/http"
	"net/http/httptest"
	"github.com/unchain1ed/webapp/model/redis"
	"github.com/gin-gonic/gin"
	"testing"
)


func TestGetLoginIdBySession(t *testing.T) {
	// Arrange ---

	// p := products.Product{ID: 123, Name: "coca cola"}
	// byteProduct, _ := json.Marshal(p)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/api/login-id",
		nil,
	)
 
	redis.NewSession(c, "loginUserIdKey","root")

	// Act ---
	common.GetLoginIdBySession(c)
 
	// Assert ---
	// var product products.Product
	// err := json.Unmarshal(response.Body.Bytes(), "root")
	assert.EqualValues(t, http.StatusOK, response.Code)
	// assert.Nil(t, err)
	// fmt.Println(product)
	// assert.EqualValues(t, uint64(123), product.ID)
}