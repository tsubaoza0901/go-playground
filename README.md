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
※test用のDB作成については以下を実行

```
mysql> CREATE DATABASE goplayground_test;
```

## 3．マイグレーションファイルの実行

① アプリケーションコンテナ内へ移動

```
$ docker exec -it go-playground bash
```

② マイグレーションファイルの実行

```
root@fe385569a625:/go/app# goose up
```
※test用のマイグレーション実行については以下を実行

```
goose up --env test
```

## 4．アプリケーションの起動

```
root@fe385569a625:/go/app# go run main.go
```

## 参考

https://bitbucket.org/liamstask/goose/

https://github.com/pressly/goose
