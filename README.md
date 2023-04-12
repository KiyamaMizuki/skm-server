# skm_server
skm_serverはメール認証用のAPIです。 
# 動作
※ローカルホストで建てた際  
①http://localhost:1323/mai  に対してPOSTでメールアドレスを投げた際、そのメールアドレスに対して認証コードを送ります。  
その後、②http://localhost:1323/token  に対してPOSTで先ほど投げたメールアドレス、認証コードを送るとトークンを発行します。  
①から②を実行する際、30分を過ぎるとtimeoverになります。  
![mail_auth](https://user-images.githubusercontent.com/44591817/97315264-e11a3f00-18ab-11eb-9054-47003b09048a.png)

# 要件
Go 1.14.6   
mysql 8.0.21  

# インストール
goのインストール
```
brew install go
```
mysqlのインストールは以下のサイトを参考にお願いします
> https://www.dbonline.jp/mysql/install/

# 実行方法
まずmysqlのユーザーを作成してください。ユーザー名、パスワードは任意です。  
以下のコマンドで作成可能です。
```
CREATE USER 'user_name'@'host_name' IDENTIFIED BY 'password'
```
その後作成したユーザーに対してallの権限を付与してください。
```
grant all on *.* to 'user_name'@'host_name'
```
最後にmysqlのDBを作成して下さい。
```
create database "DB_NAME";
```

次にメールを送信するためのgmailのアカウント情報(メールアドレス、パスワード)を確認してください。  
それらの情報を以下の場所に.envという名前で保存してください。

```
go/
 ├api/
   ├ mail_api.go
   └ .env　　　←new!!
 ├server/
   └ mail_server.go
 ├skmdb/
 │ └val/
 │ │ ├ connect.go
 │ │ ├ go.mod
 │ │ └ go.sum
 │ │  
 │ ├ authentications
 │ ├ users.go
 │ ├ go.mod
 │ └ go.sum
 │
 ├go.mod
 └ go.sum
 ```
 .envファイルの中身は以下のようにしてください。
 ```
 .envファイル
 #config SKM_project
 export DBMS=mysql　　　
 export DBUSER=”mysqlユーザー名”
 export PASS=”mysqlのパスワード”
 export PROTOCOL=tcp(127.0.0.1:3306)
 export DBNAME="データベース名"
 export TIME="parseTime=true&loc=Asia%2FTokyo"
 #mail_serverで使う値
 export MAIL_PASS=”gmailパスワード”
 export MAIL_FROM=”gmailアドレス”
 ```
 
 次にgo buildコマンドでパッケージのインストールを行ってください。
 ```api/で実行
 go build mail_api.go
 ```
 最後に実行してください。
 ```api/で実行
 go run mail_api.go
 ```

 # skm_front
 nodeの情報をweb画面で登録するサービス
 # 実行方法
 /skm_front配下でnpm run buildを実行してください。これで反映されています。