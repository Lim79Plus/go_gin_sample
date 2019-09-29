# アプリケーションの目的

- [RealWorldProject](https://github.com/gothinkster/realworld)
    - [Go Gin](https://github.com/gothinkster/golang-gin-realworld-example-app)を参照
    - ブログアプリ
- GoでのWebAPI実装方法を理解するために作成


# アーキテクチャ

- 使用言語：Go言語
    - バージョン：go 1.13 
- ライブラリ
    - FW : gin
    - ORマッパー：gorm
    - Logger: colog
- DB : MySQL 5.6


# サンプルアプリの動かし方

## DBの起動（Docker利用）

```:docker
docker run --name some-mysql_04 -e MYSQL_ROOT_PASSWORD=mysql -d -p 13306:3306 mysql:5.6 
```

## アプリの起動

```
go mod
go run *.go
```

## APIの呼び出し

※ローカルでの実行時
※sampleData直下のJsonファイル
```
// 記事一覧 
curl -X GET localhost:8080/api/articles/

// ユーザ登録_レスポンスにJWT Token
curl -X POST -H 'Content-Type:application/json' -d @newUser.json localhost:8080/api/register/

//　ログイン_レスポンスにJWT Token
curl -X POST -H 'Content-Type:application/json' -d @LoginUser.json localhost:8080/api/login/
```

JWTの動作確認

```
// ログインユーザのみ
curl -X GET -H "Authorization:  <jwt token>" "localhost:8080/api/user/" 

// 記事の投稿
curl -X POST -H "Authorization:  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjkwNjM0NzUsImlkIjoxfQ.OWJ3GwqzJcH18cCv_5cBp1w9ObOM0zV_2ok-Dmv_D2Q" -H 'Content-Type:application/json' -d @newAirticle.json localhost:8080/api/articles/

```