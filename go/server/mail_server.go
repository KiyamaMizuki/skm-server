package server

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"local.com/go-api/skmdb"
)

type Senduser struct {
	Pass string
	From string
}

// Send 指定されたメールに送信する
func Send(to string) {
	dotenv_err := godotenv.Load()
	if dotenv_err != nil {
		log.Fatal("Error loading .env file")
	}
	var s Senduser
	config_err := envconfig.Process("mail", &s)
	if config_err != nil {
		log.Fatal(config_err.Error())
	}

	auth := smtp.PlainAuth("", s.From, s.Pass, "smtp.gmail.com")
	msg :=
		[]byte("From: [琉球大学案内アプリ]登録確認メール <" + s.From + ">\r\n" +
			"To: " + to + "\r\n" + "Subject: 件名 認証コードを発行しました\r\n" + "\r\n" +
			"認証番号:" + skmdb.Gene(to) + "\r\n" + "\r\n" +
			"30分以内に認証を完了しますようお願いします。\r\n" +
			"ご本人確認のために自動的にお送りしています。\r\n" + "\r\n" +
			"このメールアドレスでの 琉球大学案内アプリ へのログインのリクエストを受け付けました。\r\n" + "\r\n" +
			to + "で登録するには上記認証番号をスマートフォンの認証番号入力画面に入力してください。" + "\r\n" +
			"*当メールに心当たりの無い場合は、誠に恐れ入りますが破棄して頂けますよう、よろしくお願い致します。\r\n" + "")

	err := smtp.SendMail("smtp.gmail.com:587", auth, s.From, []string{to}, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		return
	}
	fmt.Print("success")
}
