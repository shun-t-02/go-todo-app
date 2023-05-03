# goバージョン
FROM golang:1.20.4-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# ワーキングディレクトリの設定
WORKDIR /go-todo-app
# port
EXPOSE 8989