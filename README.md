# clean-architecture-todo-webapp-by-golang

これは[todo-webapp-by-golang](https://github.com/mz-akihiro/todo-webapp-by-golang)をクリーンアーキテクチャに基づいて再実装することで、理解をさらに深めることを目的にしたものです。

# 使用方法

※前提としてGo言語とdockerの環境が準備できているものとします

1. dockerの動作を確認後、このリポジトリに含まれているdocker-compose.ymlをしようしてdbコンテナを立ち上げて下さい。
2. コンテナの起動を確認後、以下のコマンドを順に実行して下さい。
``` bash
# Goモジュールの初期化
go mod init clean-architecture-todo-webapp-by-golang

# 立ち上げたコンテナ内のdbにテーブルを作成する
go run setup-task/task.go
go run setup-user/user.go

# 起動
go run main.go
```
3. お手元のブラウザで[http://localhost:8080/signup.html](http://localhost:8080/signup.html)に接続し、アカウント作成をして下さい。
4. 作成したアカウントでログインを行うことで、タスクの追加を行えるようになります。


__クリーンアーキテクチャについて__

* usecaseとrepository間の依存性を逆転させるためにインターフェースを介して実装を行いました。
* 他にも依存先のコードの隠蔽や今後のテスト実装なども考慮し、全てインターフェースを介す実装を行ってます。


__クライアント側の動作について__

* タスクの追加・削除は、それぞれのエンドポイントからのステータスコードが200の時のみDOM操作を行います。