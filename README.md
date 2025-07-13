# cleanarch-go
[![CI](https://github.com/ktaki8ra/cleanarch-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ktaki8ra/cleanarch-go/actions/workflows/ci.yml)

Go言語で簡単なAPIを実装してみる。   
クリーンアーキテクチャを参考にしている。

## コンパイル方法
・DBマイグレート用実行ファイルの作成   
`db-migrate`という実行ファイルが生成される
```
make dbmigrate
```
・アプリケーション実行ファイルの作成   
`cleanarch-go`という実行ファイルが生成される
```
make build
```

## テスト
・テスト実行
```
make test
```

## 実行方法
・DBマイグレート実行
```
createdb cleanarch_sample_go && ./db-migrate
```
・アプリケーション実行
```
./cleanarch-go
```

## クライアントサンプル
クライアントのサンプルリクエスト・レスポンス
### ユーザー作成 (POST `/user/create`)
・リクエスト
```
curl -X POST -H "Content-Type: application/json" -d '{"user_id":"test01", "email":"test@example.com", "password":"p4ssw0rd"}' http://localhost:8080/user/create
```
・レスポンス
```
{"status_code":200,"message":"User Created.","user_id":"test01"}
```

### ユーザー更新 (PUT `/user/update`)
・リクエスト
```
curl -X PUT -H "Content-Type: application/json" -d '{"user_id":"test01", "new_user_id":"user01", "password":"p4ssw0rd"}' http://localhost:8080/user/update
```
・レスポンス
```
{"status_code":200,"message":"User Updated.","user_id":"user01"}
```

### ユーザー取得 (GET `/user/me`)
・リクエスト
```
curl -X GET -H "Content-Type: application/json" -d '{"user_id":"user01"}' http://localhost:8080/user/me
```
・レスポンス
```
{"status_code":200,"user_id":"user01","email":"test@example.com"}
```

### ユーザー削除 (DELETE `/user/delete`)
・リクエスト
```
curl -X DELETE -H "Content-Type: application/json" -d '{"user_id":"user01", "password":"p4ssw0rd"}' http://localhost:8080/user/delete
```
・レスポンス
```
{"status_code":200,"message":"User Deleted.","user_id":"user01"}
```

## 参考文献
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- https://qiita.com/nrslib/items/a5f902c4defc83bd46b8