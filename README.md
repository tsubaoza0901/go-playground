# go-playground

# 使用方法

## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．アプリケーションの起動

```
$ docker exec -it go-playground bash
```

```
root@fe385569a625:/go/app# go run main.go
```

# google/wire
GoのDI用パッケージの一つ。

wireは以下のフローで使用することができる。 ※手順1は初回のみ

1. google/wire のインストール
    ```
    go install github.com/google/wire/cmd/wire@latest
    ```

2. 「DIしたい対象を生成する関数」を生成する関数を定義   
    - 本レポジトリではwire.goに定義された「InitializeEvent()」を指す
    - この関数が定義されたファイルは、実ビルドには不要なためBuild Constraints（今回であれば「// +build wireinject」）を設定しておく

3.  wireコマンドで、上記のファイルから「DIしたい対象を生成する関数」をジェネレート   

    ```
    $ wire
    ```
    ※ 対象の関数に対して初めて実行した場合、xxx_gen.goというファイル（今回であればwire_gen.go）が生成される。

4.  その関数を利用

## 参考

- google/wireを使ってGoでDI(dependency injection)してみる
https://www.asobou.co.jp/blog/web/google-wire

- Wire Tutorial
https://github.com/google/wire/blob/main/_tutorial/README.md