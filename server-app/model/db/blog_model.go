package db

import (
	"log"
	"fmt"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model //共通カラム
	LoginID string
	Title string
	Content string
}
// type BlogInfo struct {
// 	Title string
// 	Content string
// 	CreatedAt string
// 	UpdatedAt string
// 	DeletedAt string
// }

//送られてきたタイトルと内容をDBに登録
func Create(loginID, title, content string) (*Blog, error){
	blog := Blog{}
	blog = Blog{LoginID: loginID, Title: title, Content: content}

	result := Db.Create(&blog)
	if result.Error != nil {
		log.Println("error", result.Error);
		// エラーが発生した場合はエラーを返す
		return nil, result.Error
	}

	return &blog, nil
}

//送られてきたタイトルと内容をDBに更新
func Edit(id, loginID, title, content string) (*Blog, error){
	blog := Blog{}
	blog = Blog{LoginID: loginID, Title: title, Content: content}

	result := Db.Table("BLOGS").Where("id = ?", id).Updates(&blog)
	if result.Error != nil {
		log.Println("error", result.Error);
		// エラーが発生した場合はエラーを返す
		return nil, result.Error
	}

	return &blog, nil
}

//DBからBLOG情報を全件取得
func GetBlogOverview() ([]Blog) {
	var blogs []Blog
	
	//MySQLからuserIdに一致する構造体userを取得
	Db.Table("BLOGS").Find(&blogs)
	
	return blogs
}

// //DBからIDによる特定のBLOG情報を全件取得
// func GetBlogViewInfoById(id string) (*Blog, error) {
// 	// blog := Blog{}
// 	var blog = Blog{}
	
// 	// MySQLからIDに一致する構造体blogを取得
// 	result := Db.Table("BLOGS").Where("ID = ?", id).Find(&blog)
// 	if result.Error != nil {
// 		// エラーが発生した場合はエラーを返す
// 		return nil, result.Error
// 	}
	
// 	return &blog, nil
// }

// DBからIDによる特定のBLOG情報を取得
func GetBlogViewInfoById(id string) (*Blog, error) {
	blog := &Blog{}
fmt.Println(id)
	// MySQLからIDに一致する構造体blogを取得
	result := Db.Table("BLOGS").Where("id = ?", id).First(blog)
	if result.Error != nil {
		// エラーが発生した場合はエラーを返す
		return nil, result.Error
	}

	return blog, nil
}