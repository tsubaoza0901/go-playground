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
root@fe385569a625:/go/app# cd infrastructure/migrations/
root@fe385569a625:/go/app/infrastructure/migrations# goose mysql "root:root@tcp(db:3306)/goplayground?parseTime=true" up
```

【補足】   

`bitbucket.org/liamstask/goose/cmd/goose` を使用している場合は、コマンドの実行は root ディレクトリで良く、コマンドも以下でよかったが、

```
root@fe385569a625:/go/app# goose up
```
以下の理由から `github.com/pressly/goose/v3/cmd/goose@latest` に変更し、それに伴って使用方法に少々違いが発生

- GraphQL 導入にあたって gqlgen の latest を使用したいが、その場合 go version 1.18以上である必要がある
- しかし、go version 1.18以上では `bitbucket.org/liamstask/goose/cmd/goose` が使用できず（サポートが終了している）、開発が継続されている `github.com/pressly/goose/v3/cmd/goose@latest` に変更する必要がある

なお、`github.com/pressly/goose/v3/cmd/goose@latest` の使用方法については以下を参照   

> ・Github | pressly/goose：https://github.com/pressly/goose   
> ・Qiita | PostgreSQL+gooseでDBマイグレーションを試してみる：https://qiita.com/kishimoto828/items/179072276799c740a3eb   

## 4．アプリケーションの起動

```
root@fe385569a625:/go/app# go run main.go
```

# その他

## UTについて

### mock作成手順

1. mock化したいメソッドのインターフェースが定義されたファイルの上部に以下の記述を追加（初回のみ）

```
//go:generate mockgen -source=$GOFILE -package=mock -destination=$GOPATH/app/src/mock/$GOFILE

```

2. ターミナルで以下のコマンドを実行

```
root@fe385569a625:/go/app# go generate ./...
```

3. mockディレクトリにmockが作成または更新されていることを確認

### UT実行
ターミナルで以下のコマンドを実行

```
root@fe385569a625:/go/app# go test ./...
```

## GraphQL

### IFに変更がある場合の更新手順

1. Schema（schema.graphqls）の手動修正

2. Schemaを用いた関連ファイル（schema.resolvers.go、models_gen.go、generated.go）の自動更新

```
root@fe385569a625:/go/app# go generate ./...
```

※ ドキュメント上は以下のコマンドを毎回実行する必要はなさそうだが、なぜか上記実行前に以下のコマンドを叩かないとエラーになる、、

```
root@fe385569a625:/go/app# go get github.com/99designs/gqlgen@v0.17.19
```

3. 自動更新されない以下ファイルの修正 ※必要に応じて
- resolver.go：依存管理ファイル（はじめに依存関係を定義したら、基本的に変更することはなさそう）
- schema.resolvers.go：Schemaに対するロジックを記載するファイル。IFの変更に伴い、ロジックを変更する必要がある場合は、こちらを変更する。

### 実行サンプル

#### mutation

query

```graphql
mutation createUser($newUser: NewUser!) {
  createUser(input: $newUser)
}
```

variables

```json
{
  "newUser": {
  	"firstName": "太郎",
    "lastName": "山田",
  	"age": 40,
    "amount": 10000,
  	"email": "xxsssx@gmail.com"
	}
}
```

#### query

```graphql
query findUsers {
  users {
    name
    age
    emailAddress
    gradeName
  }
}
```


＜参考＞   
- Fusic Tech Blog | gqlgen + EchoでgolangなGraphQLサーバを作るチュートリアル   
https://tech.fusic.co.jp/posts/2020-04-12-gqlgen-echo-sample/   

# ビジネスルール

- ユーザーは新規登録時に必ず500円以上の電子マネーをチャージすること
- チャージ金額は500円以上

 ## Domain

DTOがEntityを渡すほうが良いのか、必要な情報をプリミティブな形で渡すほうが良いのかは要検討
→ DTOにEntityの情報を渡した場合は、Usecase層での変換処理がなくて済むのでUsecase層としては扱いやすい。
ただし、いろんなところでEntityをNewしたりするので、Entityの操作が多くなる点は問題

一方、DTOをプリミティブな型のフィールドで渡す場合、Usecase層でのDTO変換およびEntity変換がめんどくさそう


DTOに必要な情報があるかのバリデーションメソッドはやすのもあり？

## Usecase

関連テーブルへの保存を依頼するのもOK（正確にはUsecaseはテーブルの分け方は気にしないので関係ないですが）
ただ、エンティティが異なるものを一つのオブジェクトとしてまとめて保存を依頼するようなことはしないほうが良さそう

例）

ユーザー
 - 別テーブルに住所を保存するけど、エンティティとしては住所はセット

残高
- 別エンティティなので、ユーザー情報と一緒にRepositoryに保存依頼はしない


### input

基本的にはDTOとして、Usecaseに渡すための構造体としてのみ使う
modelに変換するメソッドはここで用意する

また、「入力値が○以上」などのビジネスロジックとしてのバリデーションはここで行う

基本的に構造体のフィールドはプリミティブな型を利用するが、特定の型に対してバリデーションのためのメソッドを定義する必要がある場合は、独自の型を定義する

→ やっぱり、domainでNewするときにチャックするほうがよさそう。でないと、似たような型が乱立してします、、

例）

```go
// TopUpAmount ...
type TopUpAmount uint

// IsMinimumAmountOrMore ...
func (a TopUpAmount) IsMinimumAmountOrMore() bool {
	return uint(a) >= balance.MinimumTopUpAmount
}
```

### output

画面表示（レスポンス）に必要なビジネスロジックとしてのメソッドを定義する

また、基本的に構造体のフィールドの型はプリミティブなものにするが、特定の型に対して画面表示のためのメソッドを定義する必要がある場合は、独自の型を定義する

ただし、そのメソッドを使用して値を変換して表示するかどうかの判断は、Presenter層やController層の責務であるためUsecase層ではそのメソッドを使用しないこと

usecase層で表示に関するロジックを持つのはやや違和感あるが、presentor層を持たない場合はしょうがない気がする

例）

```go
// Amount ...
type Amount uint

// ToJPYString ..
func (a Amount) ToJPYString() string {
	return fmt.Sprintf("%v円", a)
}
```

## adaptors/controllers

### handler

バリデーションチェックも行うがあくまで、Usecaseに渡すための必要な情報が揃っているか（requiredレベル）の確認のみに徹する
ビジネスロジックとしてのバリデーションチェックはUsecase層で行う

## Repository

トランザクションはやはり完全な隠蔽が難しそう、、
そのため、完全な隠蔽を目指すRepositoryパターンだと、少々違和感がある

そして、DataAccessの領域にエンティティが出てくるので少々扱いにくい。。

なので、Repositoryパターンではなく、Table Data Gateway（PoEAA）パターンを採用し、データベースの情報はUsecase層でエンティティに変換して使用するほうが良さそう？
https://engineering.mercari.com/blog/entry/2019-12-19-100000/
https://matope.hatenablog.com/entry/2014/05/13/193709


注）アプリケーション層からRepository層へのお願い
- 検索時のRecord Not Foundの場合はエラーを返さず、ゼロ値で戻り値を返すこと（重要）
→ Record Not Foundの場合にエラーを返すと、Record Not Foundの場合を前提に正常な処理を行いたい場合（例えば、重複チェックなど）とレコードが見つからなかった場合にエラーにしたい場合のハンドリングが煩雑になるため（アプリケーション層でエラーがRecord Not Foundかどうかを判定し、エラーで上がってきているにも関わらずエラーをなしにして、正常処理を続ける必要があるなどやや面倒）

例）gormの場合
First()など対象レコードが見つからなかった場合エラーを返すメソッドではなく、Limit(1).Find()を用いることで対象レコードが見つからない場合、ゼロ値が返却される

```go
func (r BalanceRepository) fetchBy(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.dbConn.Where("user_id = ?", userID).First(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchBlanceResultDTO(*balanceDBModel), nil
}
```
↓
```go
func (r BalanceRepository) fetchBy(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	balanceDBModel := new(dbModel.Balance)
	if err := r.dbConn.Where("user_id = ?", userID).Limit(1).Find(&balanceDBModel).Error; err != nil {
		return nil, err
	}
	return dbModel.MakeFetchBlanceResultDTO(*balanceDBModel), nil
}
```

Updateで対象のレコードがなかった場合はエラーを返して

## Presenter

クリーンアーキテクチャにある以下のような記述を見ると、モノリシックなアプリケーションを想定しているように感じる
昨今ではフロントエンド側でのロジックもあるので、ここで言うPresenter と Viewについては、バックエンドとしてはあまり気にしなくてよさそう

> GUI のユニットテストは難しい。なぜなら、画面に適切な要素が表示されている かを確認するテストを書くのが非常に難しいからだ。しかし、GUI の振る舞いの大部分は、簡 単にテストできる。Humble Object パターンを使えば、2 種類の振る舞いを Presenter と View の 2 つのクラスに分けられる。


# その他
errorは全レイヤー横断的に使用する。Domainとして定義（=errorの定義内容のみレイヤーを飛び越えて使用）？
Usecase層でエラーを設定することはなさそう？（アプリケーション内部のエラーは基本Domain層で決定できそう）