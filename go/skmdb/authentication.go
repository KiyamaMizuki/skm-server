package skmdb

import (
	"math/rand"
	"strconv"
	"time"

	"local.com/db-module/model"

	"local.com/db-module/val"
)


// Gene メールを格納し、コードを返す
func Gene(mail string) string {
	db := val.MysqlConnect()

	// 実行完了後DB接続を閉じる

	defer db.Close()
	//マイグレーション
	db.AutoMigrate(&model.Authentications{})

	//データベースに値を格納
	db.Create(&model.Authentications{Mailaddress: mail, Authcode: Authcode(), CreatedTime: time.Now(), DeletedTime: time.Now().Add(30 * time.Minute)})

	//SELECT
	i := model.Authentications{}
	db.Find(&i)
	return strconv.Itoa(i.Authcode)
}

//Gettime 時間を返す関数
func Gettime(mail string, authcode string) time.Time {
	db := val.MysqlConnect()

	// 実行完了後DB接続を閉じる
	defer db.Close()
	auth := model.Authentications{}
	num, _ := strconv.Atoi(authcode)
	db.Where("Mailaddress = ? AND Authcode = ?", mail, num).Find(&auth)
	return auth.DeletedTime
}

//Getmail メールアドレスを返す関数
func Getmail(mail string, authcode string) string {
	db := val.MysqlConnect()

	// 実行完了後DB接続を閉じる
	defer db.Close()
	auth := model.Authentications{}
	num, _ := strconv.Atoi(authcode)
	db.Where("Mailaddress = ? AND Authcode = ?", mail, num).Find(&auth)
	return auth.Mailaddress
}

//Authcode は四桁の乱数を返します。
func Authcode() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(8999)
	n = n + 1000
	return int(n)
}
