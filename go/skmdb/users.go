package skmdb

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"local.com/db-module/model"

	"local.com/db-module/val"
)

// Entry 仮DBに情報が入ってる場合のみ本番DBに登録
func Entry(mail string, authcode string) string {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	//マイグレーション
	db.AutoMigrate(&model.Users{})

	//auth := Authentication{}
	use := model.Users{}
	var token string

	if Getmail(mail, authcode) == mail {
		if Gettime(mail, authcode).After(time.Now()) {
			db.Create(&model.Users{Mailaddress: mail, Token: GenerateRandomString(), Username: getname(mail), CreatedTIME: time.Now(), DeletedTIME: time.Now()})
			db.Where("Mailaddress = ?", mail).Find(&use)
			token = use.Token
		} else {
			token = "timeover"
		}
	} else {
		token = "error"
	}
	return token
}

//Judge関数　取得したmailとDBにmailの判定
func Judge(mail string, authcode string) bool {
	db := val.MysqlConnect()
	var anc bool
	anc = false
	// 実行完了後DB接続を閉じる
	defer db.Close()
	if Getmail(mail, authcode) == mail {
		if Gettime(mail, authcode).After(time.Now()) {
			anc = true
		}
	}
	return anc
}

//ユーザーネーム取得
func getname(mailadd string) string {
	userlen := strings.Index(mailadd, "@")
	return mailadd[:userlen]
}

//GenerateRandomBytes 配列にランダムな数値を格納
func GenerateRandomBytes() []byte {
	c := 21
	b := make([]byte, c)
	_, err := rand.Read(b)

	if err != nil {
	}
	return b
}

// GenerateRandomString ランダムに生成した数値をエンコード
func GenerateRandomString() string {
	b := GenerateRandomBytes()
	return base64.URLEncoding.EncodeToString(b)
}

//GetUsername Tokenからuser名取得
func GetUsername(token string) string {
	db := val.MysqlConnect()
	// 実行完了後DB接続を閉じる
	defer db.Close()
	//マイグレーション
	db.AutoMigrate(&model.Users{})

	user := model.Users{}
	var username string
	db.Where("Token = ?",token).First(&user)
	username = user.Username
	if(username == ""){
		return "not found"
	}
	return username
	
}