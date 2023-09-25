// login_test.go

package login

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"github.com/unchain1ed/webapp/controller/login"
	"github.com/unchain1ed/webapp/model/entity"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// モックを作成するための構造体を定義
type MockedDB struct {
    mock.Mock
}

// MockedRedis を定義
type MockedRedis struct {
    mock.Mock
}


// GetOneUser メソッドを定義
func (m *MockedDB) GetOneUser(userId string) (entity.User, error) {
    args := m.Called(userId)
    return args.Get(0).(entity.User), args.Error(1)
}

// CheckUser メソッドを定義
func (m *MockedDB) CheckUser(userId string, password string) (entity.User, error) {
    args := m.Called(userId, password)
    return args.Get(0).(entity.User), args.Error(1)
}


// NewSession メソッドを定義
func (m *MockedRedis) NewSession(c *gin.Context, key string, value interface{}) error {
    args := m.Called(c, key, value)
    return args.Error(0)
}

// GetSession メソッドを定義
func (m *MockedRedis) GetSession(c *gin.Context, key string) (interface{}, error) {
    args := m.Called(c, key)
    return args.Get(0), args.Error(1)
}


func TestGetLogin(t *testing.T) {
    // モックを作成
    mockDB := new(MockedDB)
    mockRedis := new(MockedRedis)

    //環境変数設定
	//main.goからの相対パス指定
	envErr := godotenv.Load("./.env")
    if envErr != nil {
        fmt.Println("Error loading .env file", envErr)
    }

    // テストケース内で環境変数を設定
    os.Setenv("MYSQL_USER", "root")
    os.Setenv("MYSQL_PASSWORD", "password")
    os.Setenv("MYSQL_DATABASE", "user_info")
    os.Setenv("MYSQL_LOCAL_HOST", "localhost:3306")

  

    // テスト用のHTTPリクエストとレスポンスを作成
    router := gin.Default()




    // テスト用のセッションクッキーを設定
    sessionCookie := &http.Cookie{
        Name:  "loginUserIdKey",     // セッションのクッキー名を指定
        Value: "sample_session",   // セッションの値を指定
        Path:  "/",                // クッキーが適用されるパス
    }

    // レスポンスを記録するレコーダーを作成
    w := httptest.NewRecorder()

    // リクエストを作成
    req, _ := http.NewRequest("GET", "/login", nil)
    req.AddCookie(sessionCookie) // リクエストにセッションクッキーを追加

    // リクエストを処理
    router.ServeHTTP(w, req)

      // モックが呼び出されることを期待
    // user := entity.User{ID:1,Name:"root",}
    user := entity.User{
        ID:   1,
        UserId: "root",
    }
   
    mockRedis.On("GetSession", mock.Anything, mock.Anything).Return("sample_session", nil)
    mockDB.On("GetOneUser", mock.Anything).Return(user, nil)

    router.GET("/login", func(c *gin.Context) {login.GetLogin(c)})

    // ステータスコードが期待通りであることを確認
    assert.Equal(t, http.StatusOK, w.Code)

    // モックが正しく呼び出されたことを検証
    mockDB.AssertExpectations(t)
    mockRedis.AssertExpectations(t)

    // テストが終了したら環境変数をクリアする
    os.Unsetenv("MYSQL_PASSWORD")

}

func TestPostLogin(t *testing.T) {
    // モックを作成
    mockDB := new(MockedDB)
    mockRedis := new(MockedRedis)

    // モックが呼び出されることを期待
    mockDB.On("CheckUser", mock.Anything, mock.Anything).Return(entity.User{}, nil)
    mockRedis.On("NewSession", mock.Anything, mock.Anything, mock.Anything).Return(nil)

    // テスト用のHTTPリクエストとレスポンスを作成
    router := gin.Default()
    router.POST("/login", func(c *gin.Context) {login.PostLogin(c)})
    req, _ := http.NewRequest("POST", "/login", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // ステータスコードが期待通りであることを確認
    assert.Equal(t, http.StatusOK, w.Code)

    // モックが正しく呼び出されたことを検証
    mockDB.AssertExpectations(t)
    mockRedis.AssertExpectations(t)
}
