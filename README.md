# go-playground

# 使用方法

## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．アプリケーションの起動

① アプリケーションコンテナ内へ移動

```
$ docker exec -it go-playground bash
```

② アプリケーションの起動

```
root@fe385569a625:/go/app# go run main.go
```
