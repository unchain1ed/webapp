package service

import(
	"errors"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-playground/validator/v10"
)

//バリデーションチェック処理
func ValidationCheck(c *gin.Context,  err error) bool {

	// JSONデータをUser構造体にバインドしてバリデーションを実行
	if err != nil {
		// バリデーションエラーが発生した場合はエラーレスポンスを返す
		var verr validator.ValidationErrors
		if ok := errors.As(err, &verr); ok {
			var errors []string
			for _, e := range verr {
				errors = append(errors, fmt.Sprintf("%s validation failed on the %s field", e.Tag(), e.Field()))
			}
			log.Println(errors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return false
		}
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}