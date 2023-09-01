FROM golang:latest

# ホストのファイルをコンテナにコピー
COPY server-app/ /server-app

# .env ファイルをコンテナ内にコピー
COPY ./server-app/build/app/.env ./server-app/build/app/.env

# RUN apk update && apk add git
WORKDIR /server-app/cmd/webapp

# ポートを開放
EXPOSE 8080

# ビルドコマンドや実行コマンド
# RUN go build -o main .

# コンテナが起動した際に実行するコマンドやアプリケーションを指定
CMD ["go", "run", "main.go"]
