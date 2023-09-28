package configs

import "os"

type SendEmailData struct {
	APIPublicKey string
	APISecretKey string
	FromEmail    string
	FromName     string
}

func GetSendEmailData() SendEmailData {
	return SendEmailData{
		APIPublicKey: os.Getenv("MAILJET_PUBLIC_KEY"),
		APISecretKey: os.Getenv("MAILJET_SECRET_KEY"),
		FromEmail:    os.Getenv("MAILJET_FROM_EMAIL"),
		FromName:     "Easy Wallet Team",
	}
}
