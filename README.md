練習のためのものだ！

docker-composeの実行方法
```
$ docker-compose up -d mysql
```

今回のデータベースの接続情報を設定する
```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=practice \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3000 \
    MYSQL_DATABASE=Go-Practice
```
