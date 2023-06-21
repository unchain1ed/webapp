package db

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model //共通カラム
	Title string
	Content string
}

// //送られてきたUserIdと一致するUserを取得
// //取得したUserの暗号化済みPasswordと送られてきたPasswordをCompareHashAndPassword関数でチェック
// //UserのLoggedInメソッドはtop.htmlで現在ログイン中かの確認に使います
// func Login(userId, password string) (*User, error) {
// 	user :=User{}
// 	//MySQLからuserIdに一致する構造体userを取得
// 	Db.Table("USERS").Where("user_id = ?",userId).First(&user)

// 		if user.ID == 0 {
// 			err := errors.New("UserIdが一致するユーザーが存在しません。")
// 			fmt.Println(err)
// 			return nil, err
// 		}

// 		//ハッシュ化したpasswordを比較
// 		compareErr := crypto.CompareHashAndPassword(user.Password, password)
	
// 		if compareErr != nil {
// 			fmt.Println("パスワードが一致しません。:", compareErr)
// 			return nil, compareErr
// 		}
		
// 		return &user, nil
// }


//送られてきたタイトルと内容をDBに登録
func Create(title, content string) (*Blog, error){
	blog := Blog{}

	blog = Blog{Title: title, Content: content}
	Db.Create(&blog)

	return &blog, nil
}

// //gormのUpdate関数で記事情報を編集
// func Update(userId, password string) (*User, error){
// 	user := User{}

// 	// Db.Table("USERS").Where("user_id = ?", userId).First(&user)


// 	// 	if user.ID == 0 {
// 	// 		err := errors.New("UserIdが一致するユーザーが存在しません。")
// 	// 		fmt.Println(err)
// 	// 		return nil, err
// 	// 	}

// 	// //ハッシュ化したpasswordを作成
// 	// encryptPw, err := crypto.PasswordEncrypt(password)

// 	// if err != nil {
// 	// 	fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
// 	// 	return nil, err
// 	// }

// 	// user = User{UserId: userId, Password: encryptPw}


// 	//指定されたフィールドのみを更新
// 	Db.Table("USERS").Where("user_id = ?", userId).Updates(&User{UserId: userId})


// 	//モデルごと更新
// 	// Db.Model(&User{}).Where("id = ?", 1).Updates(user)


// 	return &user, nil
// }

// func GetOneUser(UserId string) (User) {
// 	user := User{}
// 	//MySQLからuserIdに一致する構造体userを取得
// 	Db.Table("USERS").Where("user_id = ?", UserId).First(&user)

// 	return user
// }