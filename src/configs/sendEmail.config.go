package configs

import "easy-wallet-be/src/utils"

type SendEmailData struct {
	APIPublicKey string
	APISecretKey string
	FromEmail    string
	FromName     string
}

func GetSendEmailData() SendEmailData {
	return SendEmailData{
		APIPublicKey: utils.GetEnv("MAILJET_PUBLIC_KEY"),
		APISecretKey: utils.GetEnv("MAILJET_SECRET_KEY"),
		FromEmail:    utils.GetEnv("MAILJET_FROM_EMAIL"),
		FromName:     "Easy Wallet Team",
	}
}
