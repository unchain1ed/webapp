# webapp
 ブログ記事管理システム
 
 【機能】
 ブログ記事閲覧、新規作成、編集、消去
 会員情報新規登録、編集
 会員情報承認
 

 【画面】
・ログイン画面

・ブログ記事一覧画面
├── 記事新規作成画面
├── 記事編集画面
├── 記事消去ダイアログ

・会員情報変更画面

・ログアウト画面


 【サーバーサイド構成】
server-app
├── build
│   ├── app
│   │   ├──.env
│   │   └── Dockerfile
│   └── db
│       ├──my.cnf
│       ├── Dockerfile
│       └──data
│  			└──.env
│       └──init
│  			└──1_create.sql
├── cmd
│   └── webapp
│       └── main.go
├── controller
│   ├── common_controller.go
│   ├── delete_controller.go
│   ├── edit_controller.go
│   ├── regist_controller.go
│   ├── setting_controller.go
│   ├── login_controller.go
│   ├── router.go
│   └── home_controller.go
│   └── dto
│       └── blog_dto.go
├── crypto
│   └── crypto.go
├── certificate
│    └──  localhost.crt
│    └──  localhost.key
├── model
│   └── entity
│       └── blog_entity.go
│       └── user_dto.go
│   └── db
│       └── blog_model.go
│       └── user_model.go
│       └── database.go
│   └── redis
│       └── redis.go
├── service
│   └──  validator.go
├── docker-compose.yml
├── go.mod
├── go.sum
