# ブログ記事管理システム
 
 【機能】
 ブログ記事閲覧、新規作成、編集、消去
 会員情報新規登録、編集
 会員情報承認
 

 【画面】

・ログイン画面
<img width="1439" alt="ログイン" src="https://github.com/unchain1ed/webapp/assets/73862261/d6f7680b-2731-433b-a6e6-a67a54d5458a">

・ブログ記事一覧画面
<img width="1440" alt="概要" src="https://github.com/unchain1ed/webapp/assets/73862261/04a1bb87-c1ce-4d9c-9cc2-928dbcfab934">

├── 個別ブログ閲覧画面
<img width="1440" alt="個別ブログ" src="https://github.com/unchain1ed/webapp/assets/73862261/815fe02c-fd2f-4cf0-a79f-acdf1782266e">

├── 記事新規作成画面
<img width="1440" alt="作成" src="https://github.com/unchain1ed/webapp/assets/73862261/1fd64e27-d19d-4ab7-8212-1cd16a4aee29">

├── 記事編集画面
<img width="1440" alt="編集" src="https://github.com/unchain1ed/webapp/assets/73862261/4047657a-01a4-424b-ad59-c4d33a55013b">

├── 記事消去ダイアログ
<img width="1440" alt="消去ダイアログ" src="https://github.com/unchain1ed/webapp/assets/73862261/c4bbbbb2-402f-4492-b0e9-6fecb5f0fc8a">


・会員情報登録画面
<img width="1440" alt="登録" src="https://github.com/unchain1ed/webapp/assets/73862261/cb8582b0-ec02-42c9-9dfb-ae486305e17d">


・会員情報変更画面
<img width="1440" alt="ID編集" src="https://github.com/unchain1ed/webapp/assets/73862261/b3d3ba41-3dac-446e-a764-4fdb9f477686">


・ログアウト画面
<img width="1439" alt="ログアウト" src="https://github.com/unchain1ed/webapp/assets/73862261/b82d8cb8-b009-4d98-ad16-ae4f6a6d53d4">


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
