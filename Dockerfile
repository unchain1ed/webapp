FROM golang:latest

# ホストのファイルをコンテナにコピー
COPY server-app/ /server-app

# .env ファイルをコンテナ内にコピー
COPY server-app/build/app/.env /server-app/build/app/.env

# RUN apk update && apk add git
WORKDIR /server-app

# # ポートを開放
# ENV PORT 8080
# # ENV HOST 0.0.0.0

# EXPOSE 8080

# ビルドコマンドや実行コマンド
RUN go build -o main .

# EXPOSE $PORT

# コンテナが起動した際に実行するコマンドやアプリケーションを指定
# CMD ["go", "build", "main.go"]
# CMD ["go", "run", "main.go", "--host", "0.0.0.0", "--port", "8080"]
# CMD ["go", "run", "main.go"]
CMD ["./main"]