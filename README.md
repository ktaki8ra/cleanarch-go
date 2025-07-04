# cleanarch-go
Go言語で簡単なAPIを実装してみる。   
クリーンアーキテクチャを参考にしている。

## コンパイル方法
・DBマイグレート用実行ファイルの作成   
`db-migrate`という実行ファイルが生成される
```
$ make dbmigrate
```
・アプリケーション実行ファイルの作成   
`cleanarch-go`という実行ファイルが生成される
```
$ make build
```

## 実行方法
・DBマイグレート実行
```
$ createdb cleanarch_sample_go
$ ./db-migrate
```
・アプリケーション実行
```
$ ./cleanarch-go
```

## クライアントサンプル
・ユーザー作成 (POST `/user/create`)
```
$ curl -X POST -H "Content-Type: application/json" -d '{"user_id":"test01", "email":"test@example.com", "password":"p4ssw0rd"}' http://localhost:8080/user/create
{"status_code":200,"message":"User Created.","user_id":"test01"}
```
・ユーザー更新 (PUT `/user/update`)
```
$ curl -X PUT -H "Content-Type: application/json" -d '{"user_id":"test01", "new_user_id":"user01", "password":"p4ssw0rd"}' http://localhost:8080/user/update
{"status_code":200,"message":"User Updated.","user_id":"user01"}
```
・ユーザー取得 (GET `/user/me`)
```
$ curl -X GET -H "Content-Type: application/json" -d '{"user_id":"user01"}' http://localhost:8080/user/me
{"status_code":200,"user_id":"user01","email":"test@example.com"}
```
・ユーザー削除 (DELETE `/user/delete`)
```
$ curl -X DELETE -H "Content-Type: application/json" -d '{"user_id":"user01", "password":"p4ssw0rd"}' http://localhost:8080/user/delete
{"status_code":200,"message":"User Deleted.","user_id":"user01"}
```
