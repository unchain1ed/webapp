package db

import (
	"fmt"
	// "os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
)

var Db *gorm.DB

func init() {
	// //環境変数設定
	// envErr := godotenv.Load()
    // if envErr != nil {
    //     fmt.Println("Error loading .env file", envErr)
    // }
	// //環境変数取得
	// user := os.Getenv("MYSQL_USER")
	// pw := os.Getenv("MYSQL_PASSWORD")
	// db_name := os.Getenv("MYSQL_DATABASE")
	// var path string = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)

	//mysql接続path
	var path string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
        "root", "password", "localhost:3306", "user_info")

	dialector := mysql.Open(path)
	var err error
	if Db , err = gorm.Open(dialector); err != nil { //Db構造体に取得結果代入
		connect(dialector, 100)
	}
	fmt.Println("db connected!!")
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if Db, err = gorm.Open(dialector); err!= nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}