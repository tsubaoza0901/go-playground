# go-playground

# 使用方法

## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．データベースの作成

① DB コンテナ内へ移動

```
$ docker exec -it go-playground-db bash
```

② DB 接続

```
root@ec19d85976f4:/# mysql -u root -h db -p
Enter password:
```

③ DB 作成

```
mysql> CREATE DATABASE goplayground;
```

## 3．マイグレーションファイルの実行

① アプリケーションコンテナ内へ移動

```
$ docker exec -it go-playground bash
```

② マイグレーションファイルの実行

```
root@fe385569a625:/go/src/app/server_side# goose up
```

## 4．アプリケーションの起動

```
root@fe385569a625:/go/app# go run main.go
```

# その他
使用する際にはルート直下に「.env」ファイルを作成し、以下の内容を追加する必要あり。

```
SIGNINGKEY=
```