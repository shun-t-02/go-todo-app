# go-todo-app
go学習用アプリ

# see
https://github.com/mattn/todo-webapp

# quick start

イメージのビルド・コンテナの起動
```
docker-compose up -d --build
```

go-todo-appコンテナにアクセス
```
docker-compose exec go-todo-app sh
```

アプリケーション起動
```
go run cmd/todoapp/main.go
```