package db

import (
	"errors"
	"fmt"
	"github.com/unchain1ed/server-app/crypto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model //共通カラム
	UserId string
	Password string
}

func init() {
	//MySQLのストレージエンジンInnoDB,Userテーブル自動生成
	 Db.Table("USERS").Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(User{})
}

func (u *User) LoggedIn() bool {
	//ログアウト中ユーザーには設定されない
	return u.ID != 0
}

//送られてきたUserId,Passwordと一致するUserが既に登録されているか確認
//PasswordEncrypt関数で送られてきたPasswordを暗号化
//gormのCreate関数で新規会員登録
func Signup(userId, password string) (*User, error){
	user := User{}

	Db.Table("USERS").Where("user_id = ?", userId).First(&user)

	if user.ID != 0 {
		err := errors.New("同一名のUserIdが既に登録されています。")
		fmt.Println(err)
		return nil, err
	}

	encryptPw, err := crypto.PasswordEncrypt(password)

	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}

	user = User{UserId: userId, Password: encryptPw}
	Db.Create(&user)
	return &user, nil
}

//送られてきたUserIdと一致するUserを取得
//取得したUserの暗号化済みPasswordと送られてきたPasswordをCompareHashAndPassword関数でチェック
//UserのLoggedInメソッドはtop.htmlで現在ログイン中かの確認に使います
func Login(userId, password string) (*User, error) {
	user :=User{}
	Db.Where("user_id = ?",userId).First(&user)

		if user.ID == 0 {
			err := errors.New("UserIdが一致するユーザーが存在しません。")
			fmt.Println(err)
			return nil, err
		}

		//ハッシュ化したpasswordを比較
		compareErr := crypto.CompareHashAndPassword(user.Password, password)
	
		if compareErr != nil {
			fmt.Println("パスワードが一致しません。:", compareErr)
			return nil, compareErr
		}

		return &user, nil
}

// func Index(w http.ResponseWriter, r *http.Request) {
// 	var allArticles []Article
// 	utility.Db.Find(&allArticles)
// 	if err := tmpl.ExecuteTemplate(w, "index.html", allArticles); err != nil {
// 		log.Fatal(err)
// 	}
// }
