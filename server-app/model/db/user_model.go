package db

import (
	"log"
	"errors"
	"github.com/unchain1ed/server-app/crypto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model //共通カラム
	UserId string
	Password string
}

func init() {
	//MySQLのストレージエンジンInnoDB,テーブル自動生成
	Db.Table("BLOGS").Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(Blog{})
	Db.Table("USERS").Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(User{})
}

func (u *User) LoggedIn() *User {
	if u.ID != 0 {
	//ログアウト中ユーザーには設定されない
	return u
	}
	return nil
}

//送られてきたUserIdと一致するUserを取得
//取得したUserの暗号化済みPasswordと送られてきたPasswordをCompareHashAndPassword関数でチェック
//UserのLoggedInメソッドはtop.htmlで現在ログイン中かの確認に使います
func Login(userId, password string) (*User, error) {
	user :=User{}
	//MySQLからuserIdに一致する構造体userを取得
	Db.Table("USERS").Where("user_id = ?",userId).First(&user)

		if user.ID == 0 {
			err := errors.New("UserIdが一致するユーザーが存在しません。")
			log.Println(err)
			return nil, err
		}

		//ハッシュ化したpasswordを比較
		compareErr := crypto.CompareHashAndPassword(user.Password, password)
	
		if compareErr != nil {
			log.Println("パスワードが一致しません。:", compareErr)
			return nil, compareErr
		}
		
		return &user, nil
}

//送られてきたUserId,Passwordと一致するUserが既に登録されているか確認
//PasswordEncrypt関数で送られてきたPasswordを暗号化
//gormのCreate関数で新規会員登録
func Signup(userId, password string) (*User, error){
	user := User{}

	Db.Table("USERS").Where("user_id = ?", userId).First(&user)

	if user.ID != 0 {
		err := errors.New("同一名のUserIdが既に登録されています。")
		log.Println(err)
		return nil, err
	}

	//ハッシュ化したpasswordを作成
	encryptPw, err := crypto.PasswordEncrypt(password)

	if err != nil {
		log.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return nil, err
	}

	user = User{UserId: userId, Password: encryptPw}
	Db.Create(&user)
	return &user, nil
}

//gormのUpdate関数でID情報編集
func UpdateId(changeId string, nowId string) (*User, error){
	user := User{}

	if err := Db.Table("USERS").Where("user_id = ?", changeId).First(&user).Error; err == nil {
			// err := errors.New("UserIdが一致するユーザーが存在しません。")
			log.Println("UserIdが重複するユーザーが存在しています。")
			log.Println("Error duplicate id from DB",user.ID)
			return nil, errors.New("UserIdが重複するユーザーが存在しています。")
	} else {
		log.Println("重複するユーザーはDBに存在しません。",err)
		log.Println("Request id from client",changeId)
	}

	// 既存のIDが存在しない場合、新しいユーザーを登録
	newUser := User{
		UserId: changeId,
		// 他のフィールドも必要に応じて設定
	}

	//指定されたフィールドのみを更新
	if err := Db.Table("USERS").Where("user_id = ?", nowId).Updates(&newUser).Error; err != nil {
		err := errors.New("IDの更新に失敗しました。")
		log.Println(err)
		return nil, err
}

	//成功消去の場合、消去されたBLOG情報をログ出力
	log.Println("Success change ID in DB",newUser.UserId)	



	return &newUser, nil
}

//gormのUpdate関数でPassword編集
func UpdatePassword(userId, password string) (*User, error){
	user := User{}

	// Db.Table("USERS").Where("user_id = ?", userId).First(&user)


	// 	if user.ID == 0 {
	// 		err := errors.New("UserIdが一致するユーザーが存在しません。")
	// 		fmt.Println(err)
	// 		return nil, err
	// 	}

	// //ハッシュ化したpasswordを作成
	// encryptPw, err := crypto.PasswordEncrypt(password)

	// if err != nil {
	// 	fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
	// 	return nil, err
	// }

	// user = User{UserId: userId, Password: encryptPw}


	//指定されたフィールドのみを更新
	Db.Table("USERS").Where("user_id = ?", userId).Updates(&User{UserId: userId})


	//モデルごと更新
	// Db.Model(&User{}).Where("id = ?", 1).Updates(user)


	return &user, nil
}

func GetOneUser(UserId string) (User) {
	user := User{}
	//MySQLからuserIdに一致する構造体userを取得
	Db.Table("USERS").Where("user_id = ?", UserId).First(&user)

	return user
}